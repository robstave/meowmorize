package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/robstave/meowmorize/internal/adapters/controller"
	"github.com/robstave/meowmorize/internal/adapters/repositories"
	"github.com/robstave/meowmorize/internal/domain"
	"github.com/robstave/meowmorize/internal/domain/types"
	"github.com/robstave/meowmorize/internal/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TemplateRenderer is a custom renderer for Echo
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	// Initialize Logger
	slogger := logger.InitializeLogger()
	logger.SetLogger(slogger) // Optional: If you prefer setting a package-level logger

	// Initialize Database
	dbPath := "./meowmerize.db"
	if path := os.Getenv("DB_PATH"); path != "" {
		dbPath = path
	}
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		slogger.Error("Failed to connect to database", "error", err)
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&types.Deck{}, &types.Card{})
	if err != nil {
		slogger.Error("Failed to migrate database", "error", err)
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize Repositories
	deckRepo := repositories.NewDeckRepositorySQLite(db)
	cardRepo := repositories.NewCardRepositorySQLite(db)

	// Initialize Service
	service := domain.NewService(slogger, deckRepo, cardRepo)

	// Initialize Controller
	homeController := controller.NewHomeController(service, slogger)

	// Initialize Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Template Renderer
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = renderer

	// Routes
	e.GET("/", homeController.Home)

	e.POST("/import-deck", homeController.ImportDeck)

	// Start Server
	port := "8789"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	slogger.Info("Starting server", "port", port)
	if err := e.Start(":" + port); err != nil && err != http.ErrServerClosed {
		slogger.Error("Shutting down the server", "error", err)
		log.Fatalf("Shutting down the server: %v", err)
	}
}

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/robstave/meowmorize/docs" // Import the docs generated by Swag
	"github.com/robstave/meowmorize/internal/adapters/controller"
	"github.com/robstave/meowmorize/internal/adapters/repositories"
	"github.com/robstave/meowmorize/internal/domain"
	"github.com/robstave/meowmorize/internal/domain/types"
	"github.com/robstave/meowmorize/internal/logger"
	httpSwagger "github.com/swaggo/echo-swagger" // Import Swagger middleware
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title MeowMorize Flashcard API
// @version 1.0
// @description API documentation for the MeowMorize Flashcard App.
// @host 192.168.86.176:8789
// @BasePath /api
func main() {
	// Initialize Logger
	slogger := logger.InitializeLogger()
	logger.SetLogger(slogger) // Optional: If you prefer setting a package-level logger

	// Initialize Database
	dbPath := "./meowmerize.db"
	if path := os.Getenv("DB_PATH"); path != "" {
		dbPath = path
	}
	slogger.Info("DBPath set", "dbpath", dbPath)

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		slogger.Error("Failed to connect to database", "path", dbPath, "error", err)
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
	meowController := controller.NewMeowController(service, slogger)

	// Initialize Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORS())
	/*
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*", "http://localhost:8789", "http://192.168.86.176:8789", "http://localhost:3000"}, // Update as needed
			AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
			AllowHeaders: []string{
				"Content-Type",
				"Authorization",
			},
		}))
	*/

	// Routes
	api := e.Group("/api")
	deckGroup := api.Group("/decks")

	sessionGroup := api.Group("/sessions")

	deckGroup.GET("", meowController.GetAllDecks)
	deckGroup.GET("/default", meowController.CreateDefaultDeck)
	deckGroup.GET("/:id", meowController.GetDeckByID)
	deckGroup.POST("", meowController.CreateDeck)
	deckGroup.PUT("/:id", meowController.UpdateDeck)
	deckGroup.DELETE("/:id", meowController.DeleteDeck)

	deckGroup.POST("/import", meowController.ImportDeck)
	deckGroup.GET("/export/:id", meowController.ExportDeck)

	deckGroup.POST("/stats/:id", meowController.ClearDeckStats)

	cardGroup := api.Group("/cards")
	cardGroup.POST("/stats", meowController.UpdateCardStats)
	cardGroup.GET("/:id", meowController.GetCardByID)
	cardGroup.POST("", meowController.CreateCard)
	cardGroup.PUT("/:id", meowController.UpdateCard)
	cardGroup.DELETE("/:id", meowController.DeleteCard)

	// Session Routes
	sessionGroup.POST("/start", meowController.StartSession)
	sessionGroup.GET("/next", meowController.GetNextCard)
	sessionGroup.DELETE("/clear", meowController.ClearSession)
	sessionGroup.GET("/stats", meowController.GetSessionStats)

	// Swagger endpoint
	e.GET("/swagger/*", httpSwagger.WrapHandler)

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

// internal/domain/seeder.go
package domain

import (
	"os"

	"github.com/google/uuid"
	"github.com/robstave/meowmorize/internal/domain/types"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) SeedUser() error {
	// Read default username and password from environment variables.
	// If not provided, default to "meow"
	defaultUsername := os.Getenv("DEFAULT_USER_USERNAME")
	if defaultUsername == "" {
		defaultUsername = "meow"
	}
	defaultPassword := os.Getenv("DEFAULT_USER_PASSWORD")
	if defaultPassword == "" {
		defaultPassword = "meow"
	}

	// Check if the user already exists
	existingUser, err := s.userRepo.GetUserByUsername(defaultUsername)
	if err != nil {
		return err
	}
	if existingUser != nil {
		// User already exists; no need to seed
		return nil
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(defaultPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create the user
	user := types.User{
		ID:       uuid.New().String(),
		Username: defaultUsername,
		Password: string(hashedPassword),
	}

	if err := s.userRepo.CreateUser(user); err != nil {
		return err
	}

	s.logger.Info("Seeded initial user", "username", user.Username)
	return nil
}

// internal/domain/seeder.go
package domain

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
	"github.com/robstave/meowmorize/internal/domain/types"
)

func (s *Service) SeedUser() error {
	// Check if the user already exists
	existingUser, err := s.userRepo.GetUserByUsername("meow")
	if err != nil {
		return err
	}
	if existingUser != nil {
		// User already exists
		return nil
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("meow"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create the user
	user := types.User{
		ID:       uuid.New().String(),
		Username: "meow",
		Password: string(hashedPassword),
	}

	if err := s.userRepo.CreateUser(user); err != nil {
		return err
	}

	s.logger.Info("Seeded initial user", "username", user.Username)
	return nil
}

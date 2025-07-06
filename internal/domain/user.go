// domain/user.go
package domain

import (
	"github.com/robstave/meowmorize/internal/domain/types"
	"golang.org/x/crypto/bcrypt"
)

// Implement the methods
func (s *Service) GetUserByUsername(username string) (*types.User, error) {
	return s.flashcardRepo.GetUserByUsername(username)
}

func (s *Service) CreateUser(user types.User) error {
	return s.flashcardRepo.CreateUser(user)
}

func (s *Service) GetAllUsers() ([]types.User, error) {
	return s.flashcardRepo.GetAllUsers()
}

func (s *Service) DeleteUser(userID string) error {
	return s.flashcardRepo.DeleteUser(userID)
}

func (s *Service) UpdateUserPassword(userID string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return s.flashcardRepo.UpdateUserPassword(userID, string(hashedPassword))
}

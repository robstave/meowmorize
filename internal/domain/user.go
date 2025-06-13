// domain/user.go
package domain

import (
	"github.com/robstave/meowmorize/internal/domain/types"
	"golang.org/x/crypto/bcrypt"
)

// Implement the methods
func (s *Service) GetUserByUsername(username string) (*types.User, error) {
	return s.userRepo.GetUserByUsername(username)
}

func (s *Service) CreateUser(user types.User) error {
	return s.userRepo.CreateUser(user)
}

func (s *Service) GetAllUsers() ([]types.User, error) {
	return s.userRepo.GetAllUsers()
}

func (s *Service) DeleteUser(userID string) error {
	return s.userRepo.DeleteUser(userID)
}

func (s *Service) UpdateUserPassword(userID string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return s.userRepo.UpdateUserPassword(userID, string(hashedPassword))
}

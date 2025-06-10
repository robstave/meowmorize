// domain/user.go
package domain

import (
	"github.com/robstave/meowmorize/internal/domain/types"
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

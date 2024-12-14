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

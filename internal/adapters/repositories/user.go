// internal/adapters/repositories/user_repository.go
package repositories

import (
	"errors"

	"github.com/robstave/meowmorize/internal/domain/types"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByUsername(username string) (*types.User, error)
	CreateUser(user types.User) error
}

type UserRepositorySQLite struct {
	db *gorm.DB
}

func NewUserRepositorySQLite(db *gorm.DB) UserRepository {
	return &UserRepositorySQLite{db: db}
}

func (r *UserRepositorySQLite) GetUserByUsername(username string) (*types.User, error) {
	var user types.User
	result := r.db.Where("username = ?", username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, result.Error
}

func (r *UserRepositorySQLite) CreateUser(user types.User) error {
	return r.db.Create(&user).Error
}

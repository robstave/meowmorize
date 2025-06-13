// internal/adapters/repositories/user.go
package repositories

import (
	"errors"

	"github.com/robstave/meowmorize/internal/domain/types"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByUsername(username string) (*types.User, error)
	CreateUser(user types.User) error
	GetAllUsers() ([]types.User, error)
	DeleteUser(userID string) error
	UpdateUserPassword(userID string, password string) error
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

func (r *UserRepositorySQLite) GetAllUsers() ([]types.User, error) {
	var users []types.User
	result := r.db.Find(&users)
	return users, result.Error
}

func (r *UserRepositorySQLite) DeleteUser(userID string) error {
	return r.db.Delete(&types.User{}, "id = ?", userID).Error
}

func (r *UserRepositorySQLite) UpdateUserPassword(userID string, password string) error {
	return r.db.Model(&types.User{}).Where("id = ?", userID).Update("password", password).Error
}

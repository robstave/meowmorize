package types

import (
	"time"
)

type Deck struct {
	ID           string    `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"type:varchar(100);not null" json:"name"`
	Description  string    `gorm:"type:text" json:"description"`
	IconURL      string    `gorm:"type:varchar(255)" json:"icon_url"`
	UserID       string    `gorm:"not null" json:"user_id"`            // NEW: owner of the deck
	Cards        []Card    `gorm:"many2many:deck_cards;" json:"cards"` // Updated to many-to-many
	LastAccessed time.Time `gorm:"autoUpdateTime" json:"last_accessed"`
}

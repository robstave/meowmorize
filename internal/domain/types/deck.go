package types

import (
	"time"
)

type Deck struct {
	ID           string    `gorm:"primaryKey" json:"id"` // UUID string as the primary key
	Name         string    `gorm:"type:varchar(100);not null" json:"name"`
	Cards        []Card    `gorm:"foreignKey:DeckID" json:"cards"`
	Description  string    `gorm:"type:text" json:"description"`        // New Description field
	LastAccessed time.Time `gorm:"autoUpdateTime" json:"last_accessed"` // New LastAccessed field

}

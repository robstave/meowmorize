package types

// Deck represents a collection of flashcards
type Deck struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"type:varchar(100);not null" json:"name"`
	Cards []Card `gorm:"foreignKey:DeckID" json:"cards"`
}

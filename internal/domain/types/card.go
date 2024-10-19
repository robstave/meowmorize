package types

// Card represents a single flashcard
type Card struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	DeckID uint   `gorm:"not null" json:"deck_id"`
	Front  string `gorm:"type:text;not null" json:"front"`
	Back   string `gorm:"type:text;not null" json:"back"`
}

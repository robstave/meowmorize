package types

type Card struct {
	ID     string `gorm:"primaryKey" json:"id"`
	DeckID string `gorm:"not null" json:"deck_id"`
	Front  struct {
		Text string `gorm:"type:text;not null" json:"text"`
	} `gorm:"embedded;embeddedPrefix:front_" json:"front"` // Embedded field for 'front' details
	Back struct {
		Text string `gorm:"type:text;not null" json:"text"`
	} `gorm:"embedded;embeddedPrefix:back_" json:"back"` // Embedded field for 'back' details
}

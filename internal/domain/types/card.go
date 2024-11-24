package types

type Card struct {
	ID     string    `gorm:"primaryKey" json:"id"`
	DeckID string    `gorm:"not null" json:"deck_id"`
	Front  CardFront `gorm:"embedded;embeddedPrefix:front_" json:"front"`
	Back   CardBack  `gorm:"embedded;embeddedPrefix:back_" json:"back"`
}

type CardFront struct {
	Text string `gorm:"type:text;not null" json:"text"`
}

type CardBack struct {
	Text string `gorm:"type:text;not null" json:"text"`
}

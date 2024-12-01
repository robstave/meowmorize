package types

import "time"

type Card struct {
	ID         string    `gorm:"primaryKey" json:"id"`
	DeckID     string    `gorm:"not null" json:"deck_id"`
	Front      CardFront `gorm:"embedded;embeddedPrefix:front_" json:"front"`
	Back       CardBack  `gorm:"embedded;embeddedPrefix:back_" json:"back"`
	Link       string    `gorm:"type:text" json:"link"`
	PassCount  int       `gorm:"default:0" json:"pass_count"`
	FailCount  int       `gorm:"default:0" json:"fail_count"`
	SkipCount  int       `gorm:"default:0" json:"skip_count"`
	StarRating int       `gorm:"default:0" json:"star_rating"`
	Retired    bool      `gorm:"default:false" json:"retired"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	ReviewedAt time.Time `json:"reviewed_at"`
}

type CardFront struct {
	Text string `gorm:"type:text;not null" json:"text"`
}

type CardBack struct {
	Text string `gorm:"type:text;not null" json:"text"`
}

// Define possible actions
type CardAction string

const (
	IncrementFail CardAction = "IncrementFail"
	IncrementPass CardAction = "IncrementPass"
	IncrementSkip CardAction = "IncrementSkip"
	SetStars      CardAction = "SetStars"
	Retire        CardAction = "Retire"
	Unretire      CardAction = "Unretire"
	ResetStats    CardAction = "ResetStats"
)

// internal/domain/types/session.go
package types

import (
	"sync"
	"time"
)

// SessionMethod represents the method used to initialize a session
type SessionMethod string

const (
	RandomMethod         SessionMethod = "Random"
	FailsMethod          SessionMethod = "Fails"
	SkipsMethod          SessionMethod = "Skips"
	WorstMethod          SessionMethod = "Worst"
	StarsMethod          SessionMethod = "Stars"
	UnratedMethod        SessionMethod = "Unrated"
	AdjustedRandomMethod SessionMethod = "AdjustedRandom"
)

// SessionLog represents a log entry for a session action.
type SessionLog struct {
	ID     string `gorm:"primaryKey" json:"id"`
	DeckID string `gorm:"index;not null" json:"deck_id"`
	CardID string `json:"card_id"` // Can be empty for reshuffle

	SessionID string `gorm:"index;not null" json:"session_id"`
	UserID    string `gorm:"not null" json:"user_id"`
	// Action can be one of: "pass", "fail", "skip", "reshuffle"
	Action    string    `gorm:"type:varchar(50);not null" json:"action"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

// CardStats represents the state of a card within a session
type CardStats struct {
	CardID  string `json:"card_id"`
	Viewed  bool   `json:"viewed"`
	Skipped bool   `json:"skipped"`
	Failed  bool   `json:"failed"`
	Passed  bool   `json:"passed"`
	Stars   int    `json:"stars"`
}

// Session represents a review session for a specific deck
type Session struct {
	DeckID    string        `json:"deckId"`
	UserID    string        `json:"userId"`
	SessionID string        `json:"sessionId"`
	CardStats []CardStats   `json:"cardStats"`
	Method    SessionMethod `json:"method"`
	Index     int           `json:"index"`
	mu        sync.Mutex    `json:"-"` // To handle concurrent access, not exported to JSON

	Stats SessionStats `json:"stats"`
}

// SessionStats holds statistics for a session
type SessionStats struct {
	TotalCards   int         `json:"totalCards"`
	ViewedCount  int         `json:"viewedCount"`
	Remaining    int         `json:"remaining"`
	CurrentIndex int         `json:"currentIndex"`
	CardStats    []CardStats `json:"cardStats"`
}

// SessionOverview represents a summary of a session
type SessionOverview struct {
	SessionID       string    `json:"sessionid"`
	DeckID          string    `json:"deck"`
	Percentage      float64   `json:"percentage"`
	PercentageAfter float64   `json:"percentage_after"`
	Cards           int       `json:"cards"`
	CardsAfter      int       `json:"cards_after"`
	Timestamp       time.Time `json:"timestamp"`
}

// GetNextCard returns the ID of the next card in the session
func (s *Session) GetNextCard() string {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.CardStats) == 0 {
		return ""
	}
	if s.Index >= len(s.CardStats) {
		s.Index = 0 // Restart the session
		resortCards(s)
	}

	cardID := s.CardStats[s.Index].CardID
	s.Index++

	// Update session stats
	s.Stats.ViewedCount++
	s.Stats.Remaining = len(s.CardStats) - s.Stats.ViewedCount
	s.Stats.CurrentIndex = s.Index

	return cardID
}

func (s *Session) GetSessionStats() SessionStats {
	s.mu.Lock()
	defer s.mu.Unlock()

	viewedCount := 0
	for _, cs := range s.CardStats {
		if cs.Viewed {
			viewedCount++
		}
	}

	stats := SessionStats{
		TotalCards:   len(s.CardStats),
		ViewedCount:  viewedCount,
		Remaining:    len(s.CardStats) - viewedCount,
		CurrentIndex: s.Index,
		CardStats:    s.CardStats,
	}

	return stats
}

// resortCards reorders cards so that skipped cards appear at the top,
// followed by failed cards, and then all remaining cards.
func resortCards(s *Session) {
	skippedCards := []CardStats{}
	failedCards := []CardStats{}
	remainingCards := []CardStats{}

	for _, cs := range s.CardStats {
		if cs.Skipped {
			skippedCards = append(skippedCards, cs)
		} else if cs.Failed {
			failedCards = append(failedCards, cs)
		} else {
			remainingCards = append(remainingCards, cs)
		}
	}

	s.CardStats = append(skippedCards, append(failedCards, remainingCards...)...)

}

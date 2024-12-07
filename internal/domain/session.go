// internal/domain/service.go
package domain

import (
	"errors"
	"math/rand"
	"sort"
	"time"

	"github.com/robstave/meowmorize/internal/domain/types"
)

// StartSession initializes or resets a session for a given deck
func (s *Service) StartSession(deckID string, count int, method types.SessionMethod) error {
	s.sessionsMu.Lock()
	defer s.sessionsMu.Unlock()

	// Fetch the deck
	deck, err := s.deckRepo.GetDeckByID(deckID)
	if err != nil {
		s.logger.Error("Failed to fetch deck", "deck_id", deckID, "error", err)
		return err
	}

	// Determine the number of cards
	totalCards := len(deck.Cards)
	if count == -1 || count > totalCards {
		count = totalCards
	}

	// Select cards based on the method
	selectedCards, err := selectCards(deck.Cards, count, method)
	if err != nil {
		s.logger.Error("Failed to select cards for session", "error", err)
		return err
	}

	deck_len := len(selectedCards)
	// Initialize CardStats
	cardStats := make([]types.CardStats, deck_len)
	for i, card := range selectedCards {
		cardStats[i] = types.CardStats{
			CardID:  card.ID,
			Viewed:  false,
			Skipped: false,
		}
	}

	stats := types.SessionStats{
		TotalCards:   deck_len,
		ViewedCount:  0,
		Remaining:    deck_len,
		CurrentIndex: 0,
	}

	// Initialize the session
	session := &types.Session{
		DeckID:    deckID,
		CardStats: cardStats,
		Method:    method,
		Index:     0,
		Stats:     stats,
	}

	// Add or reset the session in the map
	s.sessions[deckID] = session
	s.logger.Info("Session started", "deck_id", deckID) //, "method", method, "card_count", len(selectedCards))
	return nil
}

// selectCards selects cards based on the provided method
func selectCards(cards []types.Card, count int, method types.SessionMethod) ([]types.Card, error) {
	switch method {
	case types.RandomMethod:
		return selectRandomCards(cards, count), nil
	case types.FailsMethod:
		return selectTopCards(cards, count, func(c types.Card) int { return c.FailCount }), nil
	case types.SkipsMethod:
		return selectTopCards(cards, count, func(c types.Card) int { return c.SkipCount }), nil
	case types.WorstMethod:
		return selectWorstCards(cards, count), nil
	default:
		return nil, errors.New("invalid session method")
	}
}

// selectRandomCards selects random cards from the deck
func selectRandomCards(cards []types.Card, count int) []types.Card {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
	return cards[:count]
}

// selectTopCards selects top N cards based on a scoring function
func selectTopCards(cards []types.Card, count int, scoreFunc func(types.Card) int) []types.Card {
	sort.Slice(cards, func(i, j int) bool {
		return scoreFunc(cards[i]) > scoreFunc(cards[j])
	})
	if count > len(cards) {
		count = len(cards)
	}
	return cards[:count]
}

// selectWorstCards selects top N cards based on combined fail and skip counts
func selectWorstCards(cards []types.Card, count int) []types.Card {
	sort.Slice(cards, func(i, j int) bool {
		return (cards[i].FailCount + cards[i].SkipCount) > (cards[j].FailCount + cards[j].SkipCount)
	})
	if count > len(cards) {
		count = len(cards)
	}
	return cards[:count]
}

// AdjustSession updates the session based on card actions
// Current, we increment the index when we set the stats.
func (s *Service) AdjustSession(deckID string, cardID string, action types.CardAction) error {
	s.sessionsMu.RLock()
	session, exists := s.sessions[deckID]
	s.sessionsMu.RUnlock()

	// its a best effort for now
	if !exists {
		return nil
	}

	s.sessionsMu.Lock()
	s.sessionsMu.Unlock()

	// Find the card in the session
	var cardStat *types.CardStats
	for i := range session.CardStats {
		if session.CardStats[i].CardID == cardID {
			cardStat = &session.CardStats[i]
			break
		}
	}

	if cardStat == nil {
		return errors.New("card not found in session")
	}

	// Update card stats based on action
	switch action {
	case types.IncrementFail:
		cardStat.Viewed = true
		cardStat.Skipped = false
		// Optionally, update the card's fail count in the repository
	case types.IncrementPass:
		cardStat.Viewed = true
		cardStat.Skipped = false
		// Optionally, update the card's pass count in the repository
	case types.IncrementSkip:
		cardStat.Viewed = false
		cardStat.Skipped = true
	case types.SetStars:
		// none
	case types.Retire:
		cardStat.Viewed = true
		cardStat.Skipped = false
		// Optionally, retire the card in the repository
	case types.Unretire:
		cardStat.Viewed = false
		cardStat.Skipped = false
		// Optionally, unretire the card in the repository
	case types.ResetStats:
		cardStat.Viewed = false
		cardStat.Skipped = false
		// Optionally, reset stats in the repository
	default:
		return errors.New("invalid card action")
	}

	// Increment the session index
	//session.Index++
	//if session.Index >= len(session.CardStats) {
	//	session.Index = 0 // Restart the session
	//}

	// recalculate stats
	session.Stats.TotalCards = len(session.CardStats)

	var viewed = 0

	for i := range session.CardStats {
		if session.CardStats[i].Viewed {
			viewed++
		}

	}

	session.Stats.ViewedCount = viewed
	session.Stats.Remaining = session.Stats.TotalCards - viewed
	session.Stats.CurrentIndex = session.Index

	s.logger.Info("Session adjusted", "deck_id", deckID, "card_id", cardID, "action", action)
	return nil
}

// GetNextCard retrieves the next card ID in the session
// Technically, the index should already have been incremented
// so this just gets the card the index has, but perhaps I can
// add an option bool to increment too
func (s *Service) GetNextCard(deckID string) (string, error) {
	s.sessionsMu.RLock()
	session, exists := s.sessions[deckID]
	s.sessionsMu.RUnlock()

	if !exists {
		return "", errors.New("session does not exist for the given deck")
	}

	s.sessionsMu.Lock()
	defer s.sessionsMu.Unlock()

	cardID := session.GetNextCard()

	return cardID, nil
}

// ClearSession removes a session from the sessions map
func (s *Service) ClearSession(deckID string) error {
	s.sessionsMu.Lock()
	defer s.sessionsMu.Unlock()

	if _, exists := s.sessions[deckID]; !exists {
		return errors.New("session does not exist for the given deck")
	}

	delete(s.sessions, deckID)
	s.logger.Info("Session cleared", "deck_id", deckID)
	return nil
}

// GetSessionStats retrieves statistics for a given session
func (s *Service) GetSessionStats(deckID string) (types.SessionStats, error) {
	s.sessionsMu.RLock()
	session, exists := s.sessions[deckID]
	s.sessionsMu.RUnlock()

	if !exists {
		return types.SessionStats{}, nil
	}

	s.sessionsMu.Lock()
	defer s.sessionsMu.Unlock()

	stats := session.GetSessionStats()

	return stats, nil
}

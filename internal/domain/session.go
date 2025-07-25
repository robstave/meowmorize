// internal/domain/service.go
package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/robstave/meowmorize/internal/domain/types"
)

// generateSessionID generates a unique session ID
func generateSessionID() string {
	return uuid.New().String()
}

// StartSession initializes or resets a session for a given deck
func (s *Service) StartSession(deckID string, count int, method types.SessionMethod, userID string) error {
	s.sessionsMu.Lock()
	defer s.sessionsMu.Unlock()

	// Fetch the deck
	deck, err := s.deckRepo.GetDeckByID(deckID)
	if err != nil {
		s.logger.Error("Failed to fetch deck", "deck_id", deckID, "error", err)
		return err
	}

	// backfill any cards without an owner
	if err := s.backfillCardOwners(deck, userID); err != nil {
		s.logger.Error("failed to backfill card owners", "error", err)
	}

	// Update LastAccessed
	deck.LastAccessed = time.Now()
	if err := s.deckRepo.UpdateDeck(deck); err != nil {
		s.logger.Error("Failed to update deck's LastAccessed", "deck_id", deckID, "error", err)
		return err
	}
	s.logger.Info("Updated deck's LastAccessed", "deck_id", deckID, "timestamp", deck.LastAccessed)

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
			Passed:  false,
			Failed:  false,
			Stars:   card.StarRating,
		}
	}

	stats := types.SessionStats{
		TotalCards:   deck_len,
		ViewedCount:  0,
		Remaining:    deck_len,
		CurrentIndex: 0,
	}

	sessionID := generateSessionID()
	// Initialize the session
	session := &types.Session{
		DeckID:    deckID,
		UserID:    userID,
		SessionID: sessionID,
		CardStats: cardStats,
		Method:    method,
		Index:     0,
		Stats:     stats,
	}

	// Add or reset the session in the map
	s.sessions[deckID] = session
	s.logger.Info("Session started", "deck_id", deckID, "method", method, "card_count", len(selectedCards))
	return nil
}

// selectCards selects cards based on the provided method
func selectCards(cards []types.Card, count int, method types.SessionMethod) ([]types.Card, error) {
	switch method {
	case types.RandomMethod:
		return selectRandomCards(cards, count), nil
	case types.FailsMethod:
		return selectFailsCards(cards, count), nil
	case types.SkipsMethod:
		return selectSkipsCards(cards, count), nil
	case types.WorstMethod:
		return selectWorstCards(cards, count), nil
	case types.StarsMethod:
		return selectStarsCards(cards, count), nil // New Stars method
	case types.UnratedMethod:
		return selectUnratedCards(cards, count), nil // New Unrated method
	case types.AdjustedRandomMethod:
		return selectAdjustedRandomCards(cards, count), nil // New Unrated method
	default:
		return nil, errors.New("invalid session method")
	}
}

// AdjustSession updates the session based on card actions
func (s *Service) AdjustSession(deckID string, cardID string, action types.CardAction, value int, userID string) error {
	s.sessionsMu.RLock()
	session, exists := s.sessions[deckID]
	s.sessionsMu.RUnlock()

	// Best effort: if session doesn't exist, do nothing
	if !exists {
		return nil
	}

	s.sessionsMu.Lock()
	defer s.sessionsMu.Unlock()

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

	logSessionStat := false

	// Update card stats based on action
	switch action {
	case types.IncrementFail:
		cardStat.Viewed = true
		cardStat.Skipped = false
		cardStat.Failed = true
		cardStat.Passed = false
		logSessionStat = true

	case types.IncrementPass:
		cardStat.Viewed = true
		cardStat.Skipped = false
		cardStat.Failed = false
		cardStat.Passed = true
		logSessionStat = true

	case types.IncrementSkip:
		cardStat.Viewed = true
		cardStat.Skipped = true
		cardStat.Failed = false
		cardStat.Passed = false
		logSessionStat = true

	case types.SetStars:

		cardStat.Stars = value

	case types.Retire:
		cardStat.Viewed = true
		cardStat.Skipped = false
		cardStat.Failed = false
		cardStat.Passed = false
		// Implement retire logic in the repository
	case types.Unretire:
		cardStat.Viewed = false
		cardStat.Skipped = false
		cardStat.Failed = false
		cardStat.Passed = false
		// Implement unretire logic in the repository
	case types.ResetStats:
		cardStat.Viewed = false
		cardStat.Skipped = false
		cardStat.Failed = false
		cardStat.Passed = false

	default:
		return errors.New("invalid card action")
	}

	// Recalculate session stats
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

	// Hook: Log the card action.
	// (Assuming for this example that sessionID is the same as deckID and username is derived from context.)
	if logSessionStat {
		err := s.LogSessionAction(deckID, cardID, session.SessionID, userID, string(action))

		if err != nil {
			s.logger.Error("Failed to log session action", "card_id", cardID, "action", action, "error", err)
			// Optionally, you could decide to return the error.
		}
	}

	s.logger.Info("Session adjusted", "deck_id", deckID, "card_id", cardID, "action", action, "session_id", session.SessionID)
	return nil
}

// GetNextCard retrieves the next card ID in the session
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

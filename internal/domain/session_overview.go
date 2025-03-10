package domain

import (
	"github.com/robstave/meowmorize/internal/domain/types"
)

// GetSessionOverview retrieves the last 3 session overviews for a user and deck
func (s *Service) GetSessionOverview(userID string, deckID string) ([]types.SessionOverview, error) {
	// Get recent session IDs
	sessionIDs, err := s.sessionLogRepo.GetSessionLogIdsByUser(userID, deckID)
	if err != nil {
		s.logger.Error("Failed to retrieve session IDs", "user_id", userID, "deck_id", deckID, "error", err)
		return nil, err
	}

	var overviews []types.SessionOverview
	for i, sessionID := range sessionIDs {
		if i >= 3 {
			break
		}

		logs, err := s.sessionLogRepo.GetSessionLogsBySessionID(sessionID)
		if err != nil {
			s.logger.Error("Failed to retrieve session logs", "session_id", sessionID, "error", err)
			continue
		}

		if len(logs) == 0 {
			continue
		}

		overview := calculateSessionOverview(logs, sessionID, deckID)
		overview.Timestamp = logs[0].CreatedAt

		// append to the list
		overviews = append(overviews, overview)
	}

	return overviews, nil
}

// calculateSessionOverview calculates the session overview based on the logs
// for a given session and deck
// All attempts are considered, including reshuffles
// the attempts are a list in a hashmap.
// first pass is calculated by checking the metrics on the first items in the list
// final pass is calculated by checking the metrics on the last items in the list.  There may only be one item in the list.
func calculateSessionOverview(logs []types.SessionLog, sessionID, deckID string) types.SessionOverview {
	cardAttempts := map[string][]string{}

	totalFlips := 0
	for _, log := range logs {
		if log.Action == "reshuffle" {
			continue
		}
		cardAttempts[log.CardID] = append(cardAttempts[log.CardID], log.Action)
		totalFlips++
	}

	var finalPasses, initialPasses int
	totalCards := len(cardAttempts)

	for _, attempts := range cardAttempts {
		if len(attempts) == 0 {
			continue
		}

		// first attempt metric
		if attemptsFirst := attempts[0]; attemptsFirst == string(types.IncrementPass) {
			initialPasses++
		}

		// final attempt
		finalAttempt := attempts[len(attempts)-1]
		if finalAttempt == string(types.IncrementPass) {
			finalPasses++
		}

	}

	initialPercentage := (float64(initialPasses) / float64(totalCards)) * 100
	finalPercentage := (float64(finalPasses) / float64(totalCards)) * 100

	return types.SessionOverview{
		DeckID:          deckID,
		SessionID:       sessionID,
		Timestamp:       logs[0].CreatedAt,
		Cards:           totalCards,
		Percentage:      initialPercentage,
		CardsAfter:      totalFlips,
		PercentageAfter: finalPercentage,
	}
}

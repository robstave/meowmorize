package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/robstave/meowmorize/internal/domain/types"
)

// LogSessionAction logs an action for a session.
// Valid actions include: "pass", "fail", "skip", "reshuffle".
func (s *Service) LogSessionAction(deckID, cardID, sessionID, userID, action string) error {
	logEntry := types.SessionLog{
		ID:        uuid.New().String(),
		DeckID:    deckID,
		CardID:    cardID,
		SessionID: sessionID,
		UserID:    userID,
		Action:    action,
		CreatedAt: time.Now(),
	}
	if err := s.sessionLogRepo.CreateLog(logEntry); err != nil {
		s.logger.Error("Failed to log session action", "error", err)
		return err
	}
	s.logger.Info("Session action logged", "deck_id", deckID, "session_id", sessionID, "user_id", userID, "action", action)
	return nil
}

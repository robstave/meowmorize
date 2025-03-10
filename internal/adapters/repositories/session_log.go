package repositories

import (
	"fmt"

	"github.com/robstave/meowmorize/internal/domain/types"
	"gorm.io/gorm"
)

const maxSessionLogRows = 50000

// SessionLogRepository defines methods to work with session logs.
type SessionLogRepository interface {
	// CreateLog creates a new session log entry and prunes old entries if needed.
	CreateLog(log types.SessionLog) error
	// PruneLogs ensures that the total number of log entries does not exceed maxRows.
	PruneLogs(maxRows int) error

	GetSessionLogsBySessionID(sessionID string) ([]types.SessionLog, error)
	GetSessionLogIdsByUser(userID, deckID string) ([]string, error)
}

// SessionLogRepositorySQLite implements SessionLogRepository using SQLite.
type SessionLogRepositorySQLite struct {
	db *gorm.DB
}

// NewSessionLogRepositorySQLite creates a new instance of SessionLogRepositorySQLite.
func NewSessionLogRepositorySQLite(db *gorm.DB) SessionLogRepository {
	return &SessionLogRepositorySQLite{db: db}
}

// CreateLog inserts a new session log entry and prunes old entries if needed.
func (r *SessionLogRepositorySQLite) CreateLog(log types.SessionLog) error {
	if err := r.db.Create(&log).Error; err != nil {
		return err
	}
	// Prune old log entries if exceeding the limit.
	return r.PruneLogs(maxSessionLogRows)
}

// PruneLogs deletes the oldest session log entries if total rows exceed maxRows.
func (r *SessionLogRepositorySQLite) PruneLogs(maxRows int) error {
	var count int64
	if err := r.db.Model(&types.SessionLog{}).Count(&count).Error; err != nil {
		return err
	}
	if count <= int64(maxRows) {
		return nil
	}
	// Calculate the number of rows to delete.
	rowsToDelete := count - int64(maxRows)
	// Delete the oldest entries (ordered by CreatedAt ascending).
	sql := fmt.Sprintf(
		"DELETE FROM session_logs WHERE id IN (SELECT id FROM session_logs ORDER BY created_at ASC LIMIT %d)",
		rowsToDelete,
	)
	return r.db.Exec(sql).Error
}

// GetSessionLogsBySessionID retrieves all session logs for a specific session, ordered by CreatedAt ascending.
func (r *SessionLogRepositorySQLite) GetSessionLogsBySessionID(sessionID string) ([]types.SessionLog, error) {
	var logs []types.SessionLog
	err := r.db.Where("session_id = ?", sessionID).Order("created_at ASC").Find(&logs).Error
	if err != nil {
		return nil, err
	}
	return logs, nil
}

func (r *SessionLogRepositorySQLite) GetSessionLogIdsByUser(userID, deckID string) ([]string, error) {
	var sessionIDs []string

	// Direct query using window functions to get distinct session_ids ordered by their latest created_at
	query := r.db.Table("session_logs").
		Select("DISTINCT session_id").
		Where("user_id = ?", userID)

	if deckID != "" {
		query = query.Where("deck_id = ?", deckID)
	}

	err := query.
		Order("MAX(created_at) OVER (PARTITION BY session_id) DESC").
		Limit(3).
		Pluck("session_id", &sessionIDs).Error

	if err != nil {
		return nil, err
	}

	return sessionIDs, nil
}

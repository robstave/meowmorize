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

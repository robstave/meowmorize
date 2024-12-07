// internal/domain/service_session_test.go

package domain

import (
	"errors"
	"testing"

	"sync"

	"github.com/google/uuid"
	"github.com/robstave/meowmorize/internal/adapters/repositories/mocks"
	"github.com/robstave/meowmorize/internal/domain/types"
	"github.com/robstave/meowmorize/internal/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Helper function to create a new Service with mocked repositories
func setupService(deckRepo *mocks.DeckRepository, cardRepo *mocks.CardRepository) *Service {

	return &Service{
		logger:     logger.InitializeLogger(),
		deckRepo:   deckRepo,
		cardRepo:   cardRepo,
		sessions:   make(map[string]*types.Session),
		sessionsMu: sync.RWMutex{},
	}
}

func TestStartSession_Success(t *testing.T) {
	// Arrange
	deckID := uuid.New().String()
	card1 := types.Card{ID: uuid.New().String(), DeckID: deckID, Front: types.CardFront{Text: "Q1"}, Back: types.CardBack{Text: "A1"}}
	card2 := types.Card{ID: uuid.New().String(), DeckID: deckID, Front: types.CardFront{Text: "Q2"}, Back: types.CardBack{Text: "A2"}}

	deckRepoMock := &mocks.DeckRepository{}
	cardRepoMock := &mocks.CardRepository{}

	deckRepoMock.On("GetDeckByID", deckID).Return(types.Deck{
		ID:    deckID,
		Name:  "Test Deck",
		Cards: []types.Card{card1, card2},
	}, nil)

	service := setupService(deckRepoMock, cardRepoMock)

	// Act
	err := service.StartSession(deckID, 2, types.RandomMethod)

	// Assert
	assert.NoError(t, err)
	service.sessionsMu.RLock()
	session, exists := service.sessions[deckID]
	service.sessionsMu.RUnlock()
	assert.True(t, exists)
	assert.Equal(t, deckID, session.DeckID)
	assert.Equal(t, types.RandomMethod, session.Method)
	assert.Equal(t, 2, len(session.CardStats))
	assert.Equal(t, 2, session.Stats.TotalCards)
	assert.Equal(t, 0, session.Stats.ViewedCount)
	assert.Equal(t, 2, session.Stats.Remaining)
	assert.Equal(t, 0, session.Stats.CurrentIndex)
}

func TestStartSession_InvalidDeckID(t *testing.T) {
	// Arrange
	deckID := uuid.New().String()

	deckRepoMock := &mocks.DeckRepository{}
	cardRepoMock := &mocks.CardRepository{}

	deckRepoMock.On("GetDeckByID", deckID).Return(types.Deck{}, errors.New("deck not found"))

	service := setupService(deckRepoMock, cardRepoMock)

	// Act
	err := service.StartSession(deckID, 2, types.RandomMethod)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "deck not found")
	service.sessionsMu.RLock()
	_, exists := service.sessions[deckID]
	service.sessionsMu.RUnlock()
	assert.False(t, exists)
}

func TestStartSession_InsufficientCards(t *testing.T) {
	// Arrange
	deckID := uuid.New().String()
	card1 := types.Card{ID: uuid.New().String(), DeckID: deckID, Front: types.CardFront{Text: "Q1"}, Back: types.CardBack{Text: "A1"}}

	deckRepoMock := &mocks.DeckRepository{}
	cardRepoMock := &mocks.CardRepository{}

	deckRepoMock.On("GetDeckByID", deckID).Return(types.Deck{
		ID:    deckID,
		Name:  "Test Deck",
		Cards: []types.Card{card1},
	}, nil)

	service := setupService(deckRepoMock, cardRepoMock)

	// Act
	err := service.StartSession(deckID, 2, types.RandomMethod)

	// its a pass.  Just set the count to the size of the deck
	// Assert
	assert.NoError(t, err)
	service.sessionsMu.RLock()
	_, exists := service.sessions[deckID]
	service.sessionsMu.RUnlock()
	assert.True(t, exists)
}

func TestAdjustSession_Success(t *testing.T) {
	// Arrange
	deckID := uuid.New().String()
	cardID := uuid.New().String()
	card := types.Card{ID: cardID, DeckID: deckID, Front: types.CardFront{Text: "Q1"}, Back: types.CardBack{Text: "A1"}}

	deckRepoMock := &mocks.DeckRepository{}
	cardRepoMock := &mocks.CardRepository{}

	deckRepoMock.On("GetDeckByID", deckID).Return(types.Deck{
		ID:    deckID,
		Name:  "Test Deck",
		Cards: []types.Card{card},
	}, nil)

	service := setupService(deckRepoMock, cardRepoMock)
	err := service.StartSession(deckID, 1, types.RandomMethod)
	assert.NoError(t, err)

	// Mock UpdateCardStats
	cardRepoMock.On("GetCardByID", cardID).Return(&card, nil)
	cardRepoMock.On("UpdateCard", mock.Anything).Return(nil)

	// Act
	err = service.AdjustSession(deckID, cardID, types.IncrementPass)

	// Assert
	assert.NoError(t, err)
	service.sessionsMu.RLock()
	session, exists := service.sessions[deckID]
	service.sessionsMu.RUnlock()
	assert.True(t, exists)
	assert.Equal(t, true, session.CardStats[0].Viewed)
	//assert.Equal(t, 1, session.CardStats[0].PassCount)
	assert.Equal(t, 1, session.Stats.ViewedCount)
	assert.Equal(t, 0, session.Stats.Remaining)
}

func TestAdjustSession_InvalidDeckID(t *testing.T) {
	// Arrange
	deckID := uuid.New().String()
	cardID := uuid.New().String()

	deckRepoMock := &mocks.DeckRepository{}
	cardRepoMock := &mocks.CardRepository{}

	service := setupService(deckRepoMock, cardRepoMock)

	// Act
	err := service.AdjustSession(deckID, cardID, types.IncrementPass)

	// Assert no error  if there is not a session...then so what
	assert.NoError(t, err)

}

func TestGetNextCard_Success(t *testing.T) {
	// Arrange
	deckID := uuid.New().String()
	card1 := types.Card{ID: uuid.New().String(), DeckID: deckID, Front: types.CardFront{Text: "Q1"}, Back: types.CardBack{Text: "A1"}}
	card2 := types.Card{ID: uuid.New().String(), DeckID: deckID, Front: types.CardFront{Text: "Q2"}, Back: types.CardBack{Text: "A2"}}

	deckRepoMock := &mocks.DeckRepository{}
	cardRepoMock := &mocks.CardRepository{}

	deckRepoMock.On("GetDeckByID", deckID).Return(types.Deck{
		ID:    deckID,
		Name:  "Test Deck",
		Cards: []types.Card{card1, card2},
	}, nil)

	service := setupService(deckRepoMock, cardRepoMock)

	err := service.StartSession(deckID, 2, types.RandomMethod)
	assert.NoError(t, err)

	// Act
	nextCardID, err := service.GetNextCard(deckID)

	// Assert
	assert.NoError(t, err)
	assert.Contains(t, []string{card1.ID, card2.ID}, nextCardID)
}

func TestGetNextCard_SessionExhausted(t *testing.T) {
	// Arrange
	deckID := uuid.New().String()
	cardID := uuid.New().String()
	card := types.Card{ID: cardID, DeckID: deckID, Front: types.CardFront{Text: "Q1"}, Back: types.CardBack{Text: "A1"}}

	deckRepoMock := &mocks.DeckRepository{}
	cardRepoMock := &mocks.CardRepository{}

	deckRepoMock.On("GetDeckByID", deckID).Return(types.Deck{
		ID:    deckID,
		Name:  "Test Deck",
		Cards: []types.Card{card},
	}, nil)

	service := setupService(deckRepoMock, cardRepoMock)
	err := service.StartSession(deckID, 1, types.RandomMethod)
	assert.NoError(t, err)

	// First retrieval
	nextCardID, err := service.GetNextCard(deckID)
	assert.NoError(t, err)
	assert.Equal(t, cardID, nextCardID)

	// Second retrieval should restart
	nextCardID, err = service.GetNextCard(deckID)
	assert.NoError(t, err)
	assert.Equal(t, cardID, nextCardID)

}

func TestGetNextCard_InvalidDeckID(t *testing.T) {
	// Arrange
	deckID := uuid.New().String()

	deckRepoMock := &mocks.DeckRepository{}
	cardRepoMock := &mocks.CardRepository{}

	service := setupService(deckRepoMock, cardRepoMock)

	// Act
	nextCardID, err := service.GetNextCard(deckID)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "session does not exist")
	assert.Equal(t, "", nextCardID)
}

func TestClearSession_Success(t *testing.T) {
	// Arrange
	deckID := uuid.New().String()
	card := types.Card{ID: uuid.New().String(), DeckID: deckID, Front: types.CardFront{Text: "Q1"}, Back: types.CardBack{Text: "A1"}}

	deckRepoMock := &mocks.DeckRepository{}
	cardRepoMock := &mocks.CardRepository{}

	deckRepoMock.On("GetDeckByID", deckID).Return(types.Deck{
		ID:    deckID,
		Name:  "Test Deck",
		Cards: []types.Card{card},
	}, nil)

	service := setupService(deckRepoMock, cardRepoMock)
	err := service.StartSession(deckID, 1, types.RandomMethod)
	assert.NoError(t, err)

	// Act
	err = service.ClearSession(deckID)

	// Assert
	assert.NoError(t, err)
	service.sessionsMu.RLock()
	_, exists := service.sessions[deckID]
	service.sessionsMu.RUnlock()
	assert.False(t, exists)
}

func TestClearSession_NonExistentSession(t *testing.T) {
	// Arrange
	deckID := uuid.New().String()

	deckRepoMock := &mocks.DeckRepository{}
	cardRepoMock := &mocks.CardRepository{}

	service := setupService(deckRepoMock, cardRepoMock)

	// Act
	err := service.ClearSession(deckID)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "session does not exist")
}

func TestGetSessionStats_Success(t *testing.T) {
	// Arrange
	deckID := uuid.New().String()
	card1 := types.Card{ID: uuid.New().String(), DeckID: deckID, Front: types.CardFront{Text: "Q1"}, Back: types.CardBack{Text: "A1"}}
	card2 := types.Card{ID: uuid.New().String(), DeckID: deckID, Front: types.CardFront{Text: "Q2"}, Back: types.CardBack{Text: "A2"}}

	deckRepoMock := &mocks.DeckRepository{}
	cardRepoMock := &mocks.CardRepository{}

	deckRepoMock.On("GetDeckByID", deckID).Return(types.Deck{
		ID:    deckID,
		Name:  "Test Deck",
		Cards: []types.Card{card1, card2},
	}, nil)

	service := setupService(deckRepoMock, cardRepoMock)
	err := service.StartSession(deckID, 2, types.RandomMethod)
	assert.NoError(t, err)

	// Adjust session
	cardRepoMock.On("GetCardByID", card1.ID).Return(&card1, nil)
	cardRepoMock.On("UpdateCard", mock.Anything).Return(nil)
	err = service.AdjustSession(deckID, card1.ID, types.IncrementPass)
	assert.NoError(t, err)

	// Act
	stats, err := service.GetSessionStats(deckID)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 2, stats.TotalCards)
	assert.Equal(t, 1, stats.ViewedCount)
	assert.Equal(t, 1, stats.Remaining)
	assert.Equal(t, 0, stats.CurrentIndex)
}

func TestGetSessionStats_NonExistentSession(t *testing.T) {
	// Arrange
	deckID := uuid.New().String()

	deckRepoMock := &mocks.DeckRepository{}
	cardRepoMock := &mocks.CardRepository{}

	service := setupService(deckRepoMock, cardRepoMock)

	// Act
	stats, err := service.GetSessionStats(deckID)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, types.SessionStats{}, stats)
}

func TestStartSession_ConcurrentAccess(t *testing.T) {
	// This test ensures that the Service can handle concurrent StartSession calls safely.

	// Arrange
	deckID := uuid.New().String()
	card1 := types.Card{ID: uuid.New().String(), DeckID: deckID, Front: types.CardFront{Text: "Q1"}, Back: types.CardBack{Text: "A1"}}
	card2 := types.Card{ID: uuid.New().String(), DeckID: deckID, Front: types.CardFront{Text: "Q2"}, Back: types.CardBack{Text: "A2"}}

	deckRepoMock := &mocks.DeckRepository{}
	cardRepoMock := &mocks.CardRepository{}

	deckRepoMock.On("GetDeckByID", deckID).Return(types.Deck{
		ID:    deckID,
		Name:  "Concurrent Deck",
		Cards: []types.Card{card1, card2},
	}, nil)

	service := setupService(deckRepoMock, cardRepoMock)

	// Act
	done := make(chan bool)
	go func() {
		err := service.StartSession(deckID, 2, types.RandomMethod)
		assert.NoError(t, err)
		done <- true
	}()

	go func() {
		err := service.StartSession(deckID, 2, types.RandomMethod)
		// Depending on implementation, this might overwrite or return an error
		// Adjust assertions based on expected behavior
		// For example, if overwriting is allowed:
		assert.NoError(t, err)
		done <- true
	}()

	<-done
	<-done

	// Assert
	service.sessionsMu.RLock()
	session, exists := service.sessions[deckID]
	service.sessionsMu.RUnlock()
	assert.True(t, exists)
	assert.Equal(t, 2, session.Stats.TotalCards)
}

/*
func TestAdjustSession_RetireCard(t *testing.T) {
	// Arrange
	deckID := uuid.New().String()
	cardID := uuid.New().String()
	card := types.Card{
		ID:      cardID,
		DeckID:  deckID,
		Front:   types.CardFront{Text: "Q1"},
		Back:    types.CardBack{Text: "A1"},
		Retired: false,
	}

	deckRepoMock := &mocks.DeckRepository{}
	cardRepoMock := &mocks.CardRepository{}

	deckRepoMock.On("GetDeckByID", deckID).Return(types.Deck{
		ID:    deckID,
		Name:  "Retire Deck",
		Cards: []types.Card{card},
	}, nil)

	service := setupService(deckRepoMock, cardRepoMock)
	err := service.StartSession(deckID, 1, types.RandomMethod)
	assert.NoError(t, err)

	// Mock UpdateCardStats
	cardRepoMock.On("GetCardByID", cardID).Return(&card, nil)
	cardRepoMock.On("UpdateCard", mock.Anything).Return(nil)

	// Act
	err = service.AdjustSession(deckID, cardID, types.Retire, nil)

	// Assert
	assert.NoError(t, err)
	service.sessionsMu.RLock()
	session, exists := service.sessions[deckID]
	service.sessionsMu.RUnlock()
	assert.True(t, exists)
	assert.True(t, session.CardStats[0].Retired)
}
*/

func TestStartSession_DuplicateSession(t *testing.T) {
	// This test checks the behavior when attempting to start a session that's already active.

	// Arrange
	deckID := uuid.New().String()
	card := types.Card{ID: uuid.New().String(), DeckID: deckID, Front: types.CardFront{Text: "Q1"}, Back: types.CardBack{Text: "A1"}}

	deckRepoMock := &mocks.DeckRepository{}
	cardRepoMock := &mocks.CardRepository{}

	deckRepoMock.On("GetDeckByID", deckID).Return(types.Deck{
		ID:    deckID,
		Name:  "Duplicate Session Deck",
		Cards: []types.Card{card},
	}, nil)

	service := setupService(deckRepoMock, cardRepoMock)

	// Start first session
	err := service.StartSession(deckID, 1, types.RandomMethod)
	assert.NoError(t, err)

	// Act: Attempt to start a second session
	err = service.StartSession(deckID, 1, types.RandomMethod)

	// Assert: Depending on implementation, this might overwrite or return an error
	// Adjust assertions based on expected behavior
	// For this example, we'll assume it overwrites the existing session
	assert.NoError(t, err)
	service.sessionsMu.RLock()
	session, exists := service.sessions[deckID]
	service.sessionsMu.RUnlock()
	assert.True(t, exists)
	assert.Equal(t, 1, session.Stats.TotalCards)
}

func TestStartSession_MethodFallback(t *testing.T) {
	// This test ensures that if an unsupported session method is provided, the service handles it gracefully.

	// Arrange
	deckID := uuid.New().String()
	card := types.Card{ID: uuid.New().String(), DeckID: deckID, Front: types.CardFront{Text: "Q1"}, Back: types.CardBack{Text: "A1"}}

	deckRepoMock := &mocks.DeckRepository{}
	cardRepoMock := &mocks.CardRepository{}

	deckRepoMock.On("GetDeckByID", deckID).Return(types.Deck{
		ID:    deckID,
		Name:  "Fallback Method Deck",
		Cards: []types.Card{card},
	}, nil)

	service := setupService(deckRepoMock, cardRepoMock)

	// Act: Use an unsupported session method
	err := service.StartSession(deckID, 1, "UnsupportedMethod")

	// Assert: Depending on implementation, it might default to a specific method or return an error
	// For this example, we'll assume it returns an error
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid session method")
}

func TestStartSession_SessionTimeout(t *testing.T) {
	// If your session has a timeout mechanism, test that the session expires correctly.
	// This requires your Service to support session timeouts, which isn't shown in the provided code.
	// Below is a hypothetical test case.

	// Arrange
	// deckID := uuid.New().String()
	// ... similar setup as previous tests

	// Act
	// Start session
	// Wait for timeout
	// Attempt to get next card

	// Assert
	// Ensure the session has expired and appropriate error is returned
}

func TestAdjustSession_UnknownAction(t *testing.T) {
	// Arrange
	deckID := uuid.New().String()
	cardID := uuid.New().String()
	card := types.Card{
		ID:      cardID,
		DeckID:  deckID,
		Front:   types.CardFront{Text: "Q1"},
		Back:    types.CardBack{Text: "A1"},
		Retired: false,
	}

	deckRepoMock := &mocks.DeckRepository{}
	cardRepoMock := &mocks.CardRepository{}

	deckRepoMock.On("GetDeckByID", deckID).Return(types.Deck{
		ID:    deckID,
		Name:  "Unknown Action Deck",
		Cards: []types.Card{card},
	}, nil)

	service := setupService(deckRepoMock, cardRepoMock)
	err := service.StartSession(deckID, 1, types.RandomMethod)
	assert.NoError(t, err)

	// Mock UpdateCardStats
	cardRepoMock.On("GetCardByID", cardID).Return(&card, nil)

	// Act: Use an unknown action
	err = service.AdjustSession(deckID, cardID, "UnknownAction")

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid card action")
}

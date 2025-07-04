// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	types "github.com/robstave/meowmorize/internal/domain/types"
	mock "github.com/stretchr/testify/mock"
)

// MeowDomain is an autogenerated mock type for the MeowDomain type
type MeowDomain struct {
	mock.Mock
}

// AdjustSession provides a mock function with given fields: deckID, cardID, action, value, userID
func (_m *MeowDomain) AdjustSession(deckID string, cardID string, action types.CardAction, value int, userID string) error {
	ret := _m.Called(deckID, cardID, action, value, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, types.CardAction, int, string) error); ok {
		r0 = rf(deckID, cardID, action, value, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClearDeckStats provides a mock function with given fields: deckID, clearSession, clearStats
func (_m *MeowDomain) ClearDeckStats(deckID string, clearSession bool, clearStats bool) error {
	ret := _m.Called(deckID, clearSession, clearStats)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, bool, bool) error); ok {
		r0 = rf(deckID, clearSession, clearStats)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClearSession provides a mock function with given fields: deckID
func (_m *MeowDomain) ClearSession(deckID string) error {
	ret := _m.Called(deckID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(deckID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CloneCardToDeck provides a mock function with given fields: cardID, targetDeckID
func (_m *MeowDomain) CloneCardToDeck(cardID string, targetDeckID string) (*types.Card, error) {
	ret := _m.Called(cardID, targetDeckID)

	var r0 *types.Card
	if rf, ok := ret.Get(0).(func(string, string) *types.Card); ok {
		r0 = rf(cardID, targetDeckID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Card)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(cardID, targetDeckID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CollapseDecks provides a mock function with given fields: targetDeckID, sourceDeckID
func (_m *MeowDomain) CollapseDecks(targetDeckID string, sourceDeckID string) error {
	ret := _m.Called(targetDeckID, sourceDeckID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(targetDeckID, sourceDeckID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateCard provides a mock function with given fields: card, deckID, userID
func (_m *MeowDomain) CreateCard(card types.Card, deckID string, userID string) (*types.Card, error) {
	ret := _m.Called(card, deckID, userID)

	var r0 *types.Card
	if rf, ok := ret.Get(0).(func(types.Card, string, string) *types.Card); ok {
		r0 = rf(card, deckID, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Card)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Card, string, string) error); ok {
		r1 = rf(card, deckID, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateDeck provides a mock function with given fields: deck
func (_m *MeowDomain) CreateDeck(deck types.Deck) error {
	ret := _m.Called(deck)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Deck) error); ok {
		r0 = rf(deck)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateDefaultDeck provides a mock function with given fields: defaultData, userID
func (_m *MeowDomain) CreateDefaultDeck(defaultData bool, userID string) (types.Deck, error) {
	ret := _m.Called(defaultData, userID)

	var r0 types.Deck
	if rf, ok := ret.Get(0).(func(bool, string) types.Deck); ok {
		r0 = rf(defaultData, userID)
	} else {
		r0 = ret.Get(0).(types.Deck)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bool, string) error); ok {
		r1 = rf(defaultData, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUser provides a mock function with given fields: user
func (_m *MeowDomain) CreateUser(user types.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCardByID provides a mock function with given fields: cardID
func (_m *MeowDomain) DeleteCardByID(cardID string) error {
	ret := _m.Called(cardID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(cardID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDeck provides a mock function with given fields: deckID
func (_m *MeowDomain) DeleteDeck(deckID string) error {
	ret := _m.Called(deckID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(deckID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: userID
func (_m *MeowDomain) DeleteUser(userID string) error {
	ret := _m.Called(userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExportDeck provides a mock function with given fields: deckID
func (_m *MeowDomain) ExportDeck(deckID string) (types.Deck, error) {
	ret := _m.Called(deckID)

	var r0 types.Deck
	if rf, ok := ret.Get(0).(func(string) types.Deck); ok {
		r0 = rf(deckID)
	} else {
		r0 = ret.Get(0).(types.Deck)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(deckID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllDecks provides a mock function with given fields: userID
func (_m *MeowDomain) GetAllDecks(userID string) ([]types.Deck, error) {
	ret := _m.Called(userID)

	var r0 []types.Deck
	if rf, ok := ret.Get(0).(func(string) []types.Deck); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Deck)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllUsers provides a mock function with given fields:
func (_m *MeowDomain) GetAllUsers() ([]types.User, error) {
	ret := _m.Called()

	var r0 []types.User
	if rf, ok := ret.Get(0).(func() []types.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCardByID provides a mock function with given fields: cardID
func (_m *MeowDomain) GetCardByID(cardID string) (*types.Card, error) {
	ret := _m.Called(cardID)

	var r0 *types.Card
	if rf, ok := ret.Get(0).(func(string) *types.Card); ok {
		r0 = rf(cardID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Card)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(cardID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeckByID provides a mock function with given fields: deckID
func (_m *MeowDomain) GetDeckByID(deckID string) (types.Deck, error) {
	ret := _m.Called(deckID)

	var r0 types.Deck
	if rf, ok := ret.Get(0).(func(string) types.Deck); ok {
		r0 = rf(deckID)
	} else {
		r0 = ret.Get(0).(types.Deck)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(deckID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExplanation provides a mock function with given fields: prompt
func (_m *MeowDomain) GetExplanation(prompt string) (string, error) {
	ret := _m.Called(prompt)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(prompt)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(prompt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNextCard provides a mock function with given fields: deckID
func (_m *MeowDomain) GetNextCard(deckID string) (string, error) {
	ret := _m.Called(deckID)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(deckID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(deckID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSessionLogIdsByUser provides a mock function with given fields: userID, deckID
func (_m *MeowDomain) GetSessionLogIdsByUser(userID string, deckID string) ([]string, error) {
	ret := _m.Called(userID, deckID)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string, string) []string); ok {
		r0 = rf(userID, deckID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(userID, deckID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSessionLogsBySessionID provides a mock function with given fields: sessionID
func (_m *MeowDomain) GetSessionLogsBySessionID(sessionID string) ([]types.SessionLog, error) {
	ret := _m.Called(sessionID)

	var r0 []types.SessionLog
	if rf, ok := ret.Get(0).(func(string) []types.SessionLog); ok {
		r0 = rf(sessionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.SessionLog)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(sessionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSessionOverview provides a mock function with given fields: userID, deckID
func (_m *MeowDomain) GetSessionOverview(userID string, deckID string) ([]types.SessionOverview, error) {
	ret := _m.Called(userID, deckID)

	var r0 []types.SessionOverview
	if rf, ok := ret.Get(0).(func(string, string) []types.SessionOverview); ok {
		r0 = rf(userID, deckID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.SessionOverview)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(userID, deckID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSessionStats provides a mock function with given fields: deckID
func (_m *MeowDomain) GetSessionStats(deckID string) (types.SessionStats, error) {
	ret := _m.Called(deckID)

	var r0 types.SessionStats
	if rf, ok := ret.Get(0).(func(string) types.SessionStats); ok {
		r0 = rf(deckID)
	} else {
		r0 = ret.Get(0).(types.SessionStats)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(deckID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByUsername provides a mock function with given fields: username
func (_m *MeowDomain) GetUserByUsername(username string) (*types.User, error) {
	ret := _m.Called(username)

	var r0 *types.User
	if rf, ok := ret.Get(0).(func(string) *types.User); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsLLMAvailable provides a mock function with given fields:
func (_m *MeowDomain) IsLLMAvailable() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SeedUser provides a mock function with given fields:
func (_m *MeowDomain) SeedUser() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartSession provides a mock function with given fields: deckID, count, method, userID
func (_m *MeowDomain) StartSession(deckID string, count int, method types.SessionMethod, userID string) error {
	ret := _m.Called(deckID, count, method, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int, types.SessionMethod, string) error); ok {
		r0 = rf(deckID, count, method, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateCard provides a mock function with given fields: card
func (_m *MeowDomain) UpdateCard(card types.Card) error {
	ret := _m.Called(card)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Card) error); ok {
		r0 = rf(card)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateCardStats provides a mock function with given fields: cardID, action, value, deckID, userID
func (_m *MeowDomain) UpdateCardStats(cardID string, action types.CardAction, value *int, deckID string, userID string) error {
	ret := _m.Called(cardID, action, value, deckID, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, types.CardAction, *int, string, string) error); ok {
		r0 = rf(cardID, action, value, deckID, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDeck provides a mock function with given fields: deck
func (_m *MeowDomain) UpdateDeck(deck types.Deck) error {
	ret := _m.Called(deck)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Deck) error); ok {
		r0 = rf(deck)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUserPassword provides a mock function with given fields: userID, password
func (_m *MeowDomain) UpdateUserPassword(userID string, password string) error {
	ret := _m.Called(userID, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(userID, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMeowDomain interface {
	mock.TestingT
	Cleanup(func())
}

// NewMeowDomain creates a new instance of MeowDomain. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMeowDomain(t mockConstructorTestingTNewMeowDomain) *MeowDomain {
	mock := &MeowDomain{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

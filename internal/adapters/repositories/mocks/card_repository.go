// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	types "github.com/robstave/meowmorize/internal/domain/types"
)

// CardRepository is an autogenerated mock type for the CardRepository type
type CardRepository struct {
	mock.Mock
}

// CloneCardToDeck provides a mock function with given fields: cardID, targetDeckID
func (_m *CardRepository) CloneCardToDeck(cardID string, targetDeckID string) (*types.Card, error) {
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

// CountDeckAssociations provides a mock function with given fields: cardID
func (_m *CardRepository) CountDeckAssociations(cardID string) (int, error) {
	ret := _m.Called(cardID)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(cardID)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(cardID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCard provides a mock function with given fields: card
func (_m *CardRepository) CreateCard(card types.Card) error {
	ret := _m.Called(card)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Card) error); ok {
		r0 = rf(card)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCardByID provides a mock function with given fields: cardID
func (_m *CardRepository) DeleteCardByID(cardID string) error {
	ret := _m.Called(cardID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(cardID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCardByID provides a mock function with given fields: cardID
func (_m *CardRepository) GetCardByID(cardID string) (*types.Card, error) {
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

// GetCardsByDeckID provides a mock function with given fields: deckID
func (_m *CardRepository) GetCardsByDeckID(deckID string) ([]types.Card, error) {
	ret := _m.Called(deckID)

	var r0 []types.Card
	if rf, ok := ret.Get(0).(func(string) []types.Card); ok {
		r0 = rf(deckID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Card)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(deckID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCard provides a mock function with given fields: card
func (_m *CardRepository) UpdateCard(card types.Card) error {
	ret := _m.Called(card)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Card) error); ok {
		r0 = rf(card)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewCardRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewCardRepository creates a new instance of CardRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCardRepository(t mockConstructorTestingTNewCardRepository) *CardRepository {
	mock := &CardRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

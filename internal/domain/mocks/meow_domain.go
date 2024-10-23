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

// CreateDeck provides a mock function with given fields: _a0
func (_m *MeowDomain) CreateDeck(_a0 types.Deck) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Deck) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateDefaultDeck provides a mock function with given fields:
func (_m *MeowDomain) CreateDefaultDeck() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
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

// GetAllDecks provides a mock function with given fields:
func (_m *MeowDomain) GetAllDecks() ([]types.Deck, error) {
	ret := _m.Called()

	var r0 []types.Deck
	if rf, ok := ret.Get(0).(func() []types.Deck); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Deck)
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
package domain

import (
	"errors"
	"testing"

	"github.com/robstave/meowmorize/internal/adapters/repositories/mocks"
	"github.com/robstave/meowmorize/internal/domain/types"
	"github.com/robstave/meowmorize/internal/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetExplanation(t *testing.T) {
	cardRepo, userRepo, deckRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()

	// Seed user expectation
	userRepo.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	llmRepo.On("RunPrompt", mock.Anything, "test prompt").Return("answer", nil)

	s := NewService(logger.InitializeLogger(), deckRepo, cardRepo, userRepo, sessionRepo, llmRepo)
	resp, err := s.GetExplanation("test prompt")
	assert.NoError(t, err)
	assert.Equal(t, "answer", resp)
	llmRepo.AssertExpectations(t)
}

func TestGetExplanation_Error(t *testing.T) {
	cardRepo, userRepo, deckRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()

	userRepo.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	llmRepo.On("RunPrompt", mock.Anything, "bad prompt").Return("", errors.New("fail"))

	s := NewService(logger.InitializeLogger(), deckRepo, cardRepo, userRepo, sessionRepo, llmRepo)
	resp, err := s.GetExplanation("bad prompt")
	assert.Error(t, err)
	assert.Equal(t, "", resp)
	llmRepo.AssertExpectations(t)
}

func TestIsLLMAvailable(t *testing.T) {
	cardRepo, userRepo, deckRepo, sessionRepo := setupRepositories()
	llmRepo := setupLLMRepository()

	userRepo.On("GetUserByUsername", "meow").Return(&types.User{ID: "dummy", Username: "meow"}, nil)

	// LLM not initialized
	llmRepo.On("RunPrompt", mock.Anything, "test").Return("", types.ErrLLMNotInitialized)
	s := NewService(logger.InitializeLogger(), deckRepo, cardRepo, userRepo, sessionRepo, llmRepo)
	assert.False(t, s.IsLLMAvailable())

	// LLM initialized and returns without error
	llmRepo.ExpectedCalls = nil
	llmRepo.On("RunPrompt", mock.Anything, "test").Return("ok", nil)
	assert.True(t, s.IsLLMAvailable())

	// Service with nil repo
	sNil := &Service{}
	assert.False(t, sNil.IsLLMAvailable())
}

package domain

import (
	"github.com/robstave/meowmorize/internal/adapters/repositories/mocks"
)

func setupRepositories() (*mocks.CardRepository, *mocks.UserRepository, *mocks.DeckRepository, *mocks.SessionLogRepository) {
	cardRepo := new(mocks.CardRepository)
	userRepo := new(mocks.UserRepository)
	dr := new(mocks.DeckRepository)
	sessionRepo := new(mocks.SessionLogRepository)
	return cardRepo, userRepo, dr, sessionRepo
}

func setupLLMRepository() *mocks.LLMRepository {
	llmRepo := new(mocks.LLMRepository)
	return llmRepo
}

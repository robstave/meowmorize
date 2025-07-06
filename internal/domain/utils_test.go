package domain

import (
	"github.com/robstave/meowmorize/internal/adapters/repositories/mocks"
)

func setupRepositories() (*mocks.FlashcardsRepository, *mocks.SessionLogRepository) {
	flashRepo := &mocks.FlashcardsRepository{
		DeckRepository: new(mocks.DeckRepository),
		CardRepository: new(mocks.CardRepository),
		UserRepository: new(mocks.UserRepository),
	}
	sessionRepo := new(mocks.SessionLogRepository)
	return flashRepo, sessionRepo
}

func setupLLMRepository() *mocks.LLMRepository {
	llmRepo := new(mocks.LLMRepository)
	return llmRepo
}

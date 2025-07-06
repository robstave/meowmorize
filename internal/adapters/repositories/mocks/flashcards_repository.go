package mocks

type FlashcardsRepository struct {
	*DeckRepository
	*CardRepository
	*UserRepository
}

// internal/adapters/repositories/mocks/llm_repository.go
package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type LLMRepository struct {
	mock.Mock
}

func (m *LLMRepository) RunPrompt(ctx context.Context, prompt string) (string, error) {
	args := m.Called(ctx, prompt)
	return args.String(0), args.Error(1)
}

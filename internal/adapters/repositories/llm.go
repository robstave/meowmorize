// internal/adapters/repositories/llm.go
package repositories

import (
	"context"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
)

// LLMRepository defines the interface for LLM interactions
type LLMRepository interface {
	RunPrompt(ctx context.Context, prompt string) (string, error)
}

// LLMRepositoryLangChain implements LLMRepository using LangChain
type LLMRepositoryLangChain struct {
	llm llms.LLM
}

// NewLLMRepositoryLangChain creates a new LLM repository instance
func NewLLMRepositoryLangChain(apiKey string, model string) (LLMRepository, error) {
	ctx := context.Background()
	llm, err := googleai.New(ctx,
		googleai.WithAPIKey(apiKey),
		googleai.WithDefaultModel(model))
	if err != nil {
		return nil, err
	}

	return &LLMRepositoryLangChain{
		llm: llm,
	}, nil
}

// RunPrompt sends a prompt to the LLM and returns the response
func (r *LLMRepositoryLangChain) RunPrompt(ctx context.Context, prompt string) (string, error) {
	return llms.GenerateFromSinglePrompt(ctx, r.llm, prompt)
}

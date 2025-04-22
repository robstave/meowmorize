// internal/adapters/repositories/llm.go
package repositories

import (
	"context"

	types "github.com/robstave/meowmorize/internal/domain/types"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
)

// LLMRepository defines the interface for LLM interactions
type LLMRepository interface {
	RunPrompt(ctx context.Context, prompt string) (string, error)
}

// LLMRepositoryLangChain implements LLMRepository using LangChain
type LLMRepositoryLangChain struct {
	llm     llms.LLM
	enabled bool
}

// NewLLMRepositoryLangChain creates a new LLM repository instance
func NewLLMRepositoryLangChain(apiKey string, model string) (LLMRepository, error) {

	if apiKey == "" {
		// Return a disabled repository instead of an error
		return &LLMRepositoryLangChain{
			llm:     nil,
			enabled: false,
		}, nil
	}

	ctx := context.Background()
	llm, err := googleai.New(ctx,
		googleai.WithAPIKey(apiKey),
		googleai.WithDefaultModel(model))
	if err != nil {
		return nil, err
	}

	return &LLMRepositoryLangChain{
		llm:     llm,
		enabled: true,
	}, nil
}

// RunPrompt sends a prompt to the LLM and returns the response
func (r *LLMRepositoryLangChain) RunPrompt(ctx context.Context, prompt string) (string, error) {

	if !r.enabled {
		return "", types.ErrLLMNotInitialized
	}
	return llms.GenerateFromSinglePrompt(ctx, r.llm, prompt)
}

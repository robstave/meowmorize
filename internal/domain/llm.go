// internal/domain/llm.go
package domain

import (
	"context"
)

// GetExplanation sends a prompt to the LLM service and returns the response
func (s *Service) GetExplanation(prompt string) (string, error) {
	s.logger.Info("Getting LLM explanation", "prompt_length", len(prompt))

	ctx := context.Background()
	response, err := s.llmRepo.RunPrompt(ctx, prompt)
	if err != nil {
		s.logger.Error("Failed to get LLM explanation", "error", err)
		return "", err
	}

	return response, nil
}

package service

import "github.com/wangkebin/ai-code-reviewer/models"

type Llm interface {
	init(config models.Config) error
	query(query string) (string, error)
}

type AnthropicLLM struct {
	apiKey string
	model  string
}

func (a *AnthropicLLM) init(config models.Config) error {
	a.apiKey = config.OpenAIAPIKey
	a.model = config.AI_Model
	// Initialize the LLM with the API key and model
	// This is a placeholder for actual initialization logic
	return nil
}
func (a *AnthropicLLM) query(query string) (string, error) {
	// Placeholder for querying the LLM
	// This should send the query to the LLM and return the response
	return "Response from Anthropic LLM for query: " + query, nil
}

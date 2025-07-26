package models

type Config struct {
	OpenAIAPIKey string `mapstructure:"OPENAI_API_KEY"`
	AI_URL       string `mapstructure:"AI_URL"`
	AI_Model     string `mapstructure:"AI_MODEL"`
	Debug        bool   `mapstructure:"DEBUG"`
}

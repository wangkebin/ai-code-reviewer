package service

import (
	"fmt"

	"github.com/wangkebin/ai-code-reviewer/models"

	"github.com/spf13/viper"
)

func LoadConfig(path string) (config models.Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("%v", err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("%v", err)
	}
	return
}

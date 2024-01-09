package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	TelegramBot           string
	TelegramChannel       string
	URL                   string
	NotValidSentences     []string
	NotValidNeighborhoods []string
	DelayBetweenRequests  int
}

func LoadConfigs() *Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return &Config{
		TelegramBot:           viper.GetString("telegram.api_key"),
		TelegramChannel:       viper.GetString("telegram.channel_id"),
		URL:                   viper.GetString("url"),
		NotValidSentences:     viper.GetStringSlice("filters.not_valid_sentences"),
		NotValidNeighborhoods: viper.GetStringSlice("filters.not_valid_neighborhoods"),
		DelayBetweenRequests:  viper.GetInt("delay_between_requests"),
	}
}

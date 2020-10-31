package config

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"
)

type Config struct {
	Sentry                string
	TelegramBot           string
	TelegramChannel       string
	Search                string
	NotValidSentences     []string
	MinimumPictures       int
	NotValidNeighborhoods []string
	Database              string
	ShowOnlyAgency        bool
	NotValidFloors        []string
}

func LoadConfigs() *Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		sentry.CaptureException(err)
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return &Config{
		Sentry:                viper.GetString("sentry"),
		TelegramBot:           viper.GetString("telegram.api_key"),
		TelegramChannel:       viper.GetString("telegram.channel_id"),
		Search:                viper.GetString("search"),
		NotValidSentences:     viper.GetStringSlice("filters.not_valid_sentences"),
		MinimumPictures:       viper.GetInt("filters.minimum_pictures"),
		NotValidNeighborhoods: viper.GetStringSlice("filters.not_valid_neighborhoods"),
		NotValidFloors:        viper.GetStringSlice("filters.not_valid_floors"),
		Database:              viper.GetString("database"),
		ShowOnlyAgency:        viper.GetBool("filters.show_only_agencies"),
	}
}

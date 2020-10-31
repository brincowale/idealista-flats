package main

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"idealista-flats/config"
	"idealista-flats/database"
	"idealista-flats/idealista"
	"idealista-flats/telegram"
	"time"
)

func main() {
	cfg := config.LoadConfigs()
	err := sentry.Init(sentry.ClientOptions{Dsn: cfg.Sentry})
	if err != nil {
		fmt.Println(err)
	}
	defer sentry.Flush(2 * time.Second)
	t := telegram.New(cfg.TelegramBot)
	db := database.New(cfg.Database)
	client := idealista.New()
	client.Token = client.GetToken()
	results := client.GetProperties(cfg.Search)
	for _, property := range results.Properties {
		if !db.IsNewProperty(property) || !idealista.IsValidBasicProperty(property, *cfg) {
			continue
		}
		propertyDetails := client.GetProperty(property.PropertyCode)
		if idealista.IsValidDetailedProperty(propertyDetails, *cfg) {
			err := t.SendMessage(cfg.TelegramChannel, property.URL)
			if err != nil {
				sentry.CaptureException(err)
			} else {
				db.AddProperty(property)
			}
		}
	}
}

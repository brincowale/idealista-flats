package main

import (
	"idealista-flats/config"
	"idealista-flats/filters"
	"idealista-flats/idealista"
	"idealista-flats/storage"
	"idealista-flats/telegram"
	"log/slog"
	"os"
	"time"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	cfg := config.LoadConfigs()
	t := telegram.New(cfg.TelegramBot)
	delay := time.Duration(cfg.DelayBetweenRequests) * time.Second
	ic := idealista.New(delay)
	properties, err := ic.ScrapeProperties(cfg.URL)
	if err != nil {
		slog.Error("error fetching search results", "err", err)
	}
	for _, property := range properties {
		if !storage.IsNewProperty(property) {
			slog.Info("This property already exists", "property", property)
			continue
		}
		err := ic.ScrapeAdditionalDetails(property)
		if err != nil {
			slog.Error("error scraping additional results", "err", err)
		}
		isValid, reason := filters.IsValidProperty(property, cfg)
		if isValid {
			err := t.SendMessage(cfg.TelegramChannel, property.URL)
			if err != nil {
				slog.Error(err.Error(), "channel", cfg.TelegramChannel, "propertyURL", property.URL)
			}
		} else {
			slog.Info("Not a valid property", "property", property, "reason", reason)
		}
		storage.InsertProperty(property)
	}
}

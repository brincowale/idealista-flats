package filters

import (
	"idealista-flats/config"
	"idealista-flats/idealista"
	"strings"
)

const (
	InvalidDescription  = "Invalid description"
	InvalidNeighborhood = "Invalid neighborhood"
	Ok                  = "ok"
)

func IsValidProperty(property *idealista.Property, cfg *config.Config) (bool, string) {
	if !IsValidDescription(property.Description, cfg.NotValidSentences) {
		return false, InvalidDescription
	}
	if !IsValidNeighborhood(property.Neighbourhood, cfg.NotValidNeighborhoods) {
		return false, InvalidNeighborhood
	}
	return true, Ok
}

func IsValidDescription(description string, notValidSentences []string) bool {
	for _, sentence := range notValidSentences {
		if strings.Contains(strings.ToLower(description), strings.ToLower(sentence)) {
			return false
		}
	}
	return true
}

func IsValidNeighborhood(neighborhood string, notValidNeighborhoods []string) bool {
	for _, notValidNeighborhood := range notValidNeighborhoods {
		if strings.Contains(strings.ToLower(neighborhood), strings.ToLower(notValidNeighborhood)) {
			return false
		}
	}
	return true
}

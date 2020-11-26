package idealista

import (
	"idealista-flats/config"
	"strings"
)

func IsValidBasicProperty(property Property, cfg config.Config) bool {
	if !IsValidPicturesNumber(property.NumPhotos, cfg.MinimumPictures) {
		return false
	}
	if !IsValidNeighborhood(property.Neighborhood, cfg.NotValidNeighborhoods) {
		return false
	}
	return true
}

func IsValidDetailedProperty(propertyDetails PropertyDetails, cfg config.Config) bool {
	if !IsValidFloor(propertyDetails.TranslatedTexts.FloorNumberDescription, cfg.NotValidFloors) {
		return false
	}
	if cfg.ShowOnlyAgency && !IsAgency(propertyDetails.ContactInfo.Professional) {
		return false
	}
	if !IsValidAgency(propertyDetails.ContactInfo.CommercialName, cfg.NotValidAgencies) {
		return false
	}
	if !IsValidDescription(propertyDetails.PropertyComment, cfg.NotValidSentences) {
		return false
	}
	return true
}

func IsValidPicturesNumber(pictures int, min int) bool {
	if pictures >= min {
		return true
	}
	return false
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

func IsValidFloor(floor string, notValidFloors []string) bool {
	for _, notValidFloor := range notValidFloors {
		if strings.Contains(strings.ToLower(floor), strings.ToLower(notValidFloor)) {
			return false
		}
	}
	return true
}

func IsAgency(isAgency bool) bool {
	return isAgency
}

func IsValidAgency(agency string, notValidAgencies []string) bool {
	for _, notValidAgency := range notValidAgencies {
		if strings.ToLower(agency) == strings.ToLower(notValidAgency) {
			return false
		}
	}
	return true
}

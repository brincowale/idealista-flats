package storage

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"idealista-flats/idealista"
	"time"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("flats.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&Property{})
}

type Property struct {
	Id            string `gorm:"primaryKey"`
	Title         string
	URL           string
	Price         string
	Description   string
	M2            string
	Neighbourhood string
	FirstDateSeen time.Time
	LastDateSeen  time.Time
}

func InsertProperty(propertyIdealista *idealista.Property) bool {
	dbProperty := Property{
		Id:            propertyIdealista.Id,
		Title:         propertyIdealista.Title,
		URL:           propertyIdealista.URL,
		Price:         propertyIdealista.Price,
		Description:   propertyIdealista.Description,
		M2:            propertyIdealista.M2,
		Neighbourhood: propertyIdealista.Neighbourhood,
	}
	if IsNewProperty(propertyIdealista) {
		dbProperty.FirstDateSeen = time.Now()
		dbProperty.LastDateSeen = time.Now()
	} else {
		dbProperty.LastDateSeen = time.Now()
		existingProperty := Property{}
		DB.First(&existingProperty, "id = ?", propertyIdealista.Id)
		dbProperty.FirstDateSeen = existingProperty.FirstDateSeen
	}
	err := DB.Save(&dbProperty)
	return err.Error == nil
}

func IsNewProperty(propertyIdealista *idealista.Property) bool {
	dbProperty := Property{
		Id: propertyIdealista.Id,
	}
	return DB.First(&dbProperty, "id = ?", propertyIdealista.Id).Error != nil
}

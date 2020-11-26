package database

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"idealista-flats/idealista"
)

type DB struct {
	conn *gorm.DB
}

func New(dsn string) *DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		sentry.CaptureException(err)
	}
	return &DB{
		conn: db,
	}
}

func (db DB) AddProperty(property idealista.Property) {
	db.conn.Select("id").Create(idealista.Property{
		PropertyCode: fmt.Sprintf("%s_%.0f_%.0f", property.SuggestedTexts.Title, property.Size, property.Price),
	})
}

func (db DB) IsNewProperty(property idealista.Property) bool {
	var p idealista.Property
	customId := fmt.Sprintf("%s_%.0f_%.0f", property.SuggestedTexts.Title, property.Size, property.Price)
	db.conn.Where("id = ?", customId).First(&p)
	return p.PropertyCode == ""
}

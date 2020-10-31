package database

import (
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
	db.conn.Select("id").Create(property)
}

func (db DB) IsNewProperty(property idealista.Property) bool {
	var p idealista.Property
	db.conn.First(&p, property.PropertyCode)
	return p.PropertyCode == ""
}

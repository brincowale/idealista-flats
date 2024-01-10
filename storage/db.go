package storage

import (
	"database/sql"
	"idealista-flats/idealista"
	"log"
	_ "modernc.org/sqlite"
	"time"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("sqlite", "flats.db")
	if err != nil {
		log.Fatal(err)
	}
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS properties (
		id TEXT PRIMARY KEY,
		title TEXT,
		url TEXT,
		price TEXT,
		description TEXT,
		m2 TEXT,
		neighbourhood TEXT,
		first_date_seen DATETIME,
		last_date_seen DATETIME
	);`)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertProperty(propertyIdealista *idealista.Property) bool {
	query := `
		INSERT OR REPLACE INTO properties (
			id, title, url, price, description, m2, neighbourhood, first_date_seen, last_date_seen
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);
	`
	_, err := DB.Exec(query,
		propertyIdealista.Id,
		propertyIdealista.Title,
		propertyIdealista.URL,
		propertyIdealista.Price,
		propertyIdealista.Description,
		propertyIdealista.M2,
		propertyIdealista.Neighbourhood,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		log.Println("Error inserting property:", err)
		return false
	}
	return true
}

func IsNewProperty(propertyIdealista *idealista.Property) bool {
	var count int
	row := DB.QueryRow("SELECT COUNT(*) FROM properties WHERE id = ?", propertyIdealista.Id)
	if err := row.Scan(&count); err != nil {
		log.Println("Error checking if property exists:", err)
		return false
	}
	return count == 0
}

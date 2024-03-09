package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

var err error

func InitDB() {
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventTable := `
    CREATE TABLE IF NOT EXISTS eventS (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        location TEXT NOT NULL,
        dateTime DATETIME NOT NULL,
        user_id INTEGER
    )
    `

	_, err := DB.Exec(createEventTable)

	if err != nil {
		panic("Could not create the event table")
	}
}

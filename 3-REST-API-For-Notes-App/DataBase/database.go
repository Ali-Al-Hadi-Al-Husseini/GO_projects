package database

import (
	"database/sql"
	"log"
)

var db *sql.DB

func initDB() {
	var err error

	db, err = sql.Open("sqlite3", "./notes.db")

	if err != nil {
		log.Fatal(err)
	}
	createTable := `
	CREATE TABLE IF NOT EXISTS notes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		content TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

}

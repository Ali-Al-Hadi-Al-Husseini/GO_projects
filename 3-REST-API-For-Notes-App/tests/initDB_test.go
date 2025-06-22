package tests

import (
	"os"
	"testing"

	database "github.com/Ali-Al-Hadi-Al-Husseini/GO_projects/tree/main/3-REST-API-For-Notes-App/DB_Dir"
)

func Test_initDB(t *testing.T) {
	os.Remove("notes.db")
	database.InitDB()

	// check if db exists
	_, err := os.Stat("notes.db")
	if os.IsNotExist(err) {
		t.Fatal("Databse not created")

	}

	// check if table exists
	row := database.DB.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='notes'")
	var tableName string
	if err := row.Scan(&tableName); err != nil || tableName != "notes" {
		t.Fatal("Table 'notes' was not created")
	}

	// Cleanup
	database.DB.Close()
	os.Remove("notes.db")
}

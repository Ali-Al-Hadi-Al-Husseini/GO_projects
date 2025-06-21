package tests

import (
	"os"
	"testing"

	database "github.com/Ali-Al-Hadi-Al-Husseini/GO_projects/tree/main/3-REST-API-For-Notes-App/DB_Dir"
)

func Test_initDB(t *testing.T) {
	os.Remove("notes.db")
	database.InitDB()

	_, err := os.Stat("notes.db")
	if os.IsNotExist(err) {
		t.Fatal("Databse not created")

	}

}

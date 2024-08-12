package main

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
	"os"
)

var Db sql.DB

func main() {
	createDB()

}

func createDB() {
	Db, _ := sql.Open("sqlite", "data.db")
	defer Db.Close()

	createTrip, _ := Db.Prepare("CREATE TABLE IF NOT EXISTS trips (id INTEGER PRIMARY KEY, name TEXT)")
	_, err := createTrip.Exec()
	if err != nil {
		fmt.Printf("Failed to create table: %v\n", err)
		os.Exit(1)
	}
}

/*
Copyright © 2024 Nikolas Molinari <nikolas.molinari@gmail.com>
*/
package main

import (
	"database/sql"
	"fmt"
	"github.com/Gavin152/crewcut/cmd"
	_ "modernc.org/sqlite"
	"os"
)

func main() {
	createDB()
	cmd.Execute()
}

func createDB() {
	Db, _ := sql.Open("sqlite", "data.db")
	defer Db.Close()

	createTrips, _ := Db.Prepare("CREATE TABLE IF NOT EXISTS crews (id INTEGER PRIMARY KEY, name TEXT UNIQUE NOT NULL)")
	_, err := createTrips.Exec()
	if err != nil {
		fmt.Printf("Failed to create table: %v\n", err)
		os.Exit(1)
	}

	createUsers, _ := Db.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, username TEXT, crewId INTEGER, FOREIGN KEY(crewId) REFERENCES crews (id))")
	_, err = createUsers.Exec()
	if err != nil {
		fmt.Printf("Failed to create table: %v\n", err)
		os.Exit(1)
	}
}

package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "../data.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT
	);
	
	INSERT OR REPLACE INTO tasks (id, name) VALUES (0, 'This is my first Task example for Tests :)');`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
}

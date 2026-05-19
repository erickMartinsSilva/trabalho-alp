package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var Instance *sql.DB

func Init() {
	if Instance != nil {
		return
	}

	dbPath := os.Getenv("DB_PATH")
	if(dbPath == "") {
		dbPath = "./app.db"
	}

	db, err := sql.Open("sqlite3", dbPath+"?_journal_mode=WAL&_busy_timeout=5000")
	if err != nil {
		log.Fatal(err)
		return
	}
	db.SetMaxOpenConns(1)

	if err := migrate(db); err != nil {
		log.Fatal(err)
		return
	}

	Instance = db
}
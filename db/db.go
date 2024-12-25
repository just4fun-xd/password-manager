package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}

	initScript, err := os.ReadFile("db/init.sql")
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(string(initScript)); err != nil {
		return nil, err
	}

	log.Println("База данных успешно инициализирована")
	return db, nil
}

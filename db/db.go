package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() {
	DB, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("could not connect to db.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}

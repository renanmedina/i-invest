package utils

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	initDB()
}

func initDB() {
	configs := GetConfigs()
	openedDb, err := sql.Open("postgres", configs.DB_URI)

	if err != nil {
		panic(err)
	}

	db = openedDb
}

func GetDatabase() *sql.DB {
	return db
}

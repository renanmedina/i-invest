package utils

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func GetDatabase() *sql.DB {
	configs := GetConfigs()
	connection, err := sql.Open("postgres", configs.DB_URI)

	if err != nil {
		panic(err)
	}

	return connection
}

// func MigrateDb() {
// 	connection := postgres.WithInstance(GetDatabase(), &postgres.Config{})
// 	m, err := migrate.NewWithDatabaseInstance("file:///db/migrations", "postgres", connection)
// 	m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
// }

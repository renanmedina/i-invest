package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func GetDatabase() *sql.DB {
	configs := GetConfigs()
	connectionString := fmt.Sprintf("%s:authToken=%s", configs.DB_URI, configs.DB_TOKEN)
	connection, err := sql.Open("libsql", connectionString)

	if err != nil {
		panic(err)
	}

	return connection
}

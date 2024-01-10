package utils

import (
	"github.com/surrealdb/surrealdb.go"
)

func GetDatabase() *surrealdb.DB {
	configs := GetConfigs()
	connection, err := surrealdb.New(configs.DB_URI)

	if err != nil {
		panic(err)
	}

	if _, err = connection.Signin(map[string]interface{}{
		"user": configs.DB_USERNAME,
		"pass": configs.DB_PASSWORD,
	}); err != nil {
		panic(err)
	}

	if _, err = connection.Use(configs.DB_NAMESPACE, configs.DATABASE_NAME); err != nil {
		panic(err)
	}

	return connection
}

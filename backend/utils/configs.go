package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Configs struct {
	DB_URI        string
	DB_TOKEN      string
	DB_USERNAME   string
	DB_PASSWORD   string
	DB_NAMESPACE  string
	DATABASE_NAME string
}

func GetConfigs() Configs {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	return Configs{
		DB_URI:        os.Getenv("DB_URI"),
		DB_TOKEN:      os.Getenv("DB_TOKEN"),
		DB_USERNAME:   os.Getenv("DB_USERNAME"),
		DB_PASSWORD:   os.Getenv("DB_PASSWORD"),
		DB_NAMESPACE:  os.Getenv("DB_NAMESPACE"),
		DATABASE_NAME: os.Getenv("DB_NAME"),
	}
}

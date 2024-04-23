package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Configs struct {
	DB_URI       string
	DB_HOST      string
	DB_PORT      string
	DB_TOKEN     string
	DB_USERNAME  string
	DB_PASSWORD  string
	DB_NAMESPACE string
	DB_NAME      string
}

func (c *Configs) DbConnectionInfo() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.DB_HOST, c.DB_PORT, c.DB_USERNAME, c.DB_PASSWORD, c.DB_NAME,
	)
}

func GetConfigs() Configs {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	return Configs{
		DB_URI:       os.Getenv("DB_URL"),
		DB_HOST:      os.Getenv("DB_HOST"),
		DB_PORT:      os.Getenv("DB_PORT"),
		DB_TOKEN:     os.Getenv("DB_TOKEN"),
		DB_USERNAME:  os.Getenv("DB_USERNAME"),
		DB_PASSWORD:  os.Getenv("DB_PASSWORD"),
		DB_NAMESPACE: os.Getenv("DB_NAMESPACE"),
		DB_NAME:      os.Getenv("DB_NAME"),
	}
}

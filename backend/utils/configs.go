package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	LOG_FORMAT_TEXT = "text"
	LOG_FORMAT_JSON = "json"
)

type Configs struct {
	DB_URI                              string
	DB_HOST                             string
	DB_PORT                             string
	DB_TOKEN                            string
	DB_USERNAME                         string
	DB_PASSWORD                         string
	DB_NAMESPACE                        string
	DB_NAME                             string
	AWS_REGION                          string
	AWS_ACCESS_KEY                      string
	AWS_SECRET_KEY                      string
	AWS_ANNOUNCEMENTS_FILES_BUCKET_NAME string
	LOG_FORMAT                          string
}

var loadedConfigs *Configs

func init() {
	loadedConfigs = loadConfigs()
}

func (c *Configs) DbConnectionInfo() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.DB_HOST, c.DB_PORT, c.DB_USERNAME, c.DB_PASSWORD, c.DB_NAME,
	)
}

func GetConfigs() *Configs {
	return loadedConfigs
}

func loadConfigs() *Configs {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	return &Configs{
		DB_URI:                              os.Getenv("DB_URL"),
		DB_HOST:                             os.Getenv("DB_HOST"),
		DB_PORT:                             os.Getenv("DB_PORT"),
		DB_TOKEN:                            os.Getenv("DB_TOKEN"),
		DB_USERNAME:                         os.Getenv("DB_USERNAME"),
		DB_PASSWORD:                         os.Getenv("DB_PASSWORD"),
		DB_NAMESPACE:                        os.Getenv("DB_NAMESPACE"),
		DB_NAME:                             os.Getenv("DB_NAME"),
		AWS_REGION:                          os.Getenv("AWS_REGION"),
		AWS_ACCESS_KEY:                      os.Getenv("AWS_ACCESS_KEY"),
		AWS_SECRET_KEY:                      os.Getenv("AWS_SECRET_KEY"),
		AWS_ANNOUNCEMENTS_FILES_BUCKET_NAME: os.Getenv("AWS_ANNOUNCEMENTS_FILES_BUCKET_NAME"),
		LOG_FORMAT:                          os.Getenv("LOG_FORMAT"),
	}
}

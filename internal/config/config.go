package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	SECKRET_KEY string

	DATABASE_URL string

	LOG_LEVEL string

	EMAIL_ADDRESS  string
	EMAIL_PASSWORD string
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(".env file not found")
	}
	SECKRET_KEY = os.Getenv("SECRET_KEY")
	DATABASE_URL = os.Getenv("DATABASE_URL")
	LOG_LEVEL = os.Getenv("LOG_LEVEL")
	EMAIL_ADDRESS = os.Getenv("EMAIL_ADDRESS")
	EMAIL_PASSWORD = os.Getenv("EMAIL_PASSWORD")

}

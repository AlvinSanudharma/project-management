package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	AppConfig *Config
)

type Config struct {
	AppPort          string
	DBHost           string
	DBPort           string
	DBUser           string
	DBPassword       string
	DBName           string
	JWTSecret        string
	JWTExpireMinutes string
	JWTRefreshToken  string
	JWTExpire        string
}

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found")
	}

	AppConfig = &Config{
		AppPort:    getEnv("PORT", "3030"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
	}
}

func getEnv(key string, fallback string) string {
	value, exists := os.LookupEnv(key)

	if exists {
		return value
	} else {
		return fallback
	}

}

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, fallback to environment variables")
	}
}

func GetDBDSN() string {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN not set in environment")
	}
	return dsn
}

func GetJWTKEY() string {
	jwt_key := os.Getenv("JWT_KEY")
	if jwt_key == "" {
		log.Fatal("JWT_KEY not set in environment")
	}
	return jwt_key
}

func GetAPPUSER() string {
	jwt_key := os.Getenv("APP_USER")
	if jwt_key == "" {
		log.Fatal("APP_USER not set in environment")
	}
	return jwt_key
}

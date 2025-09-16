package configs

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    DatabaseURL string
    Port        string
}

func LoadConfig() (*Config, error) {
    err := godotenv.Load()
    if err != nil {
        log.Println("Error loading .env file")
        return nil, err
    }

    return &Config{
        DatabaseURL: os.Getenv("DATABASE_URL"),
        Port:        os.Getenv("PORT"),
    }, nil
}
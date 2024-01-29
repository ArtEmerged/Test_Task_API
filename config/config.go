package config

import (
	"os"

	"github.com/joho/godotenv"
)

type ConfigDB struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type Config struct {
	Port     string
	DB       ConfigDB
	DebugMod bool
}

func InitConfig() (*Config, error) {
	err := godotenv.Load("app.env")
	if err != nil {
		return nil, err
	}

	cfgDb := ConfigDB{
		Driver:   os.Getenv("DB_DRIVER"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}

	return &Config{
		Port:     os.Getenv("APP_PORT"),
		DB:       cfgDb,
		DebugMod: os.Getenv("DEBUG_MODE") == "true",
	}, nil
}

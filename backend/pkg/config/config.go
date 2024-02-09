package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Hostname string
	Port     string
	Name     string
	User     string
	Password string
}

func Load() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	config := &Config{
		Hostname: os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	return config

}

func BuildDSN(cfg *Config) string {
	dsn := ""
	if cfg.Name != "" {
		dsn += fmt.Sprintf("dbname=%v", cfg.Name)
	}
	if cfg.Hostname != "" {
		dsn += fmt.Sprintf(" host=%v", cfg.Hostname)
	}
	if cfg.Port != "" {
		dsn += fmt.Sprintf(" port=%v", cfg.Port)
	}
	if cfg.User != "" {
		dsn += fmt.Sprintf(" user=%v", cfg.User)
	}
	if cfg.Password != "" {
		dsn += fmt.Sprintf(" password=%v", cfg.Password)
	}
	return dsn
}

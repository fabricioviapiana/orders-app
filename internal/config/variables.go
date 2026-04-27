package config

import (
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DSN string
}

type Config struct {
	DB DBConfig
}

func Load() *Config {
	_ = godotenv.Load()
	var dsn string
	if dsn = os.Getenv("DATABASE_DATA_SOURCE_NAME"); dsn == "" {
		dsn = "postgres://user:password@localhost:5432/orders_db?sslmode=disable"
	}

	return &Config{
		DB: DBConfig{
			DSN: dsn,
		},
	}
}

package config

import "os"

type Config struct {
	Port        string
	DatabaseURL string
}

func Load() *Config {

	port := os.Getenv(
		"PORT",
	)

	if port == "" {
		port = "8081"
	}

	databaseURL := os.Getenv(
		"DATABASE_URL",
	)

	if databaseURL == "" {
		databaseURL = "postgres://postgres:postgres@prism_postgres_db:5432/stations_db?sslmode=disable"
	}

	return &Config{
		Port:        port,
		DatabaseURL: databaseURL,
	}
}

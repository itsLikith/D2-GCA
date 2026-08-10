package config

import "os"

type Config struct {
	Port string

	DMEServiceURL        string
	RNAVServiceURL       string
	SimulationServiceURL string
}

func Load() *Config {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	dmeURL := os.Getenv("DME_SERVICE_URL")
	if dmeURL == "" {
		dmeURL = "http://localhost:8081"
	}

	rnavURL := os.Getenv("RNAV_SERVICE_URL")
	if rnavURL == "" {
		rnavURL = "http://localhost:8080"
	}

	simURL := os.Getenv("SIMULATION_SERVICE_URL")
	if simURL == "" {
		simURL = "http://localhost:8082"
	}

	return &Config{
		Port:                 port,
		DMEServiceURL:        dmeURL,
		RNAVServiceURL:       rnavURL,
		SimulationServiceURL: simURL,
	}
}

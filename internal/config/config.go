// Package config is global configuration for application operation.
package config

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v6"
)

const (
	DefaultBaseURL       = "http://localhost:8080"
	DefaultServerAddress = ":8080"
	DefaultDataBaseURI   = "postgres://postgres:postgres@localhost:5432/exchanges_history?sslmode=disable"
)

// Config contains app configuration.
type Config struct {
	// BaseURL - base app address
	BaseURL string `env:"BASE_URL"`
	// ServerAddress - server address
	ServerAddress string `env:"SERVER_ADDRESS"`
	DataBaseURI   string `env:"DATABASE_URI"`
}

// The function checks for the presence of a flag. f - flag values
func checkExists(f string) bool {
	return flag.Lookup(f) == nil
}

func defaultConfig() Config {
	return Config{
		BaseURL:       DefaultBaseURL,
		ServerAddress: DefaultServerAddress,
		DataBaseURI:   DefaultDataBaseURI,
	}
}

func New() *Config {
	c := defaultConfig()

	err := env.Parse(&c)
	if err != nil {
		log.Fatal(err)
	}

	if checkExists("b") {
		flag.StringVar(&c.BaseURL, "b", c.BaseURL, "BaseUrl")
	}

	if checkExists("a") {
		flag.StringVar(&c.ServerAddress, "a", c.ServerAddress, "ServerAddress")
	}

	if checkExists("d") {
		flag.StringVar(&c.DataBaseURI, "d", c.DataBaseURI, "DataBaseURI")
	}

	flag.Parse()

	return &c
}

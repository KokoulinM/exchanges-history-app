// Package configs is global configuration for application operation.
package config

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v6"
)

const (
	DefaultBaseURL       = "http://localhost:8080"
	DefaultServerAddress = ":8080"
)

// Config contains app configuration.
type Config struct {
	// BaseURL - base app address
	BaseURL string `env:"BASE_URL"`
	// ServerAddress - server address
	ServerAddress string `env:"SERVER_ADDRESS"`
}

// The function checks for the presence of a flag. f - flag values
func checkExists(f string) bool {
	return flag.Lookup(f) == nil
}

func defaultConfig() Config {
	return Config{
		BaseURL:       DefaultBaseURL,
		ServerAddress: DefaultServerAddress,
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

	flag.Parse()

	return &c
}

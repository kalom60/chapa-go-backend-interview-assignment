package config

import (
	"errors"
	"os"
	"strconv"
)

var (
	ErrInvalidPort  = errors.New("port number is invalid")
	ErrInvalidDbUrl = errors.New("db url is invalid")
)

type Config struct {
	Port  int
	DbUrl string
}

func NewConfig() (*Config, error) {
	config := &Config{}
	if err := config.loadEnv(); err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Config) loadEnv() error {
	portStr := os.Getenv("APP_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return ErrInvalidPort
	}
	c.Port = port

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		return ErrInvalidDbUrl
	}
	c.DbUrl = dbUrl

	return nil
}

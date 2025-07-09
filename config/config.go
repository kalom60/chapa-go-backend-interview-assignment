package config

import (
	"errors"
	"os"
	"strconv"
)

var (
	ErrInvalidPort           = errors.New("port number is invalid")
	ErrInvalidDbUrl          = errors.New("db url is invalid")
	ErrInvalidChapaSecretKey = errors.New("chapa secret key is invalid")
	ErrInvalidChapaBaseUrl   = errors.New("chapa base url is invalid")
	ErrInvalidRedisUrl       = errors.New("redis url is invalid")
)

type Config struct {
	Port           int
	DbUrl          string
	ChapaSecretKey string
	ChapaBaseUrl   string
	RedisUrl       string
	RedisPassword  string
	Render         string
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

	chapaSecretKey := os.Getenv("CHAPA_SECRET_KEY")
	if chapaSecretKey == "" {
		return ErrInvalidChapaSecretKey
	}
	c.ChapaSecretKey = chapaSecretKey

	chapaBaseUrl := os.Getenv("CHAPA_BASE_URL")
	if chapaBaseUrl == "" {
		return ErrInvalidChapaBaseUrl
	}
	c.ChapaBaseUrl = chapaBaseUrl

	redisUrl := os.Getenv("REDIS_URL")
	if redisUrl == "" {
		return ErrInvalidRedisUrl
	}
	c.RedisUrl = redisUrl

	redisPassword := os.Getenv("REDIS_PASSWORD")
	c.RedisPassword = redisPassword

	render := os.Getenv("RENDER")
	c.Render = render

	return nil
}

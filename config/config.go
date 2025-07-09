package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/spf13/viper"
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
	WebhookSecret  string
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

	webhookSecret := os.Getenv("WEBHOOK_SECRET_KEY")
	c.WebhookSecret = webhookSecret

	err = initViper()
	if err != nil {
		return err
	}

	return nil
}

func initViper() error {
	// viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	return err
	// }

	return nil
}

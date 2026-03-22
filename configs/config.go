package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type сonfig struct {
	Db   dbConfig
	Auth authConfig
}

type dbConfig struct {
	Dsn string
}

type authConfig struct {
	Secret string
}

func LoadConfig() *сonfig {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, using default config")
	}
	return &сonfig{
		Db: dbConfig{
			Dsn: os.Getenv("DSN"),
		},
		Auth: authConfig{
			Secret: os.Getenv("TOKEN"),
		},
	}
}

func (c *сonfig) GetAuthSecret() string {
	return c.Auth.Secret
}

func (c *сonfig) GetDbConfig() string {
	return c.Db.Dsn
}

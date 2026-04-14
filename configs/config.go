package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db   DbConfig
	Auth AuthConfig
}

type DbConfig struct {
	Dsn      string
	FilePath string
}

type AuthConfig struct {
	Secret string
}

func LoadConfig() *Config {
	// Пакет godotenv читает пары «ключ‑значение» из файла .env и устанавливает их как переменные окружения процесса (не системы в целом).
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, using default config")
	}
	return &Config{
		Db: DbConfig{
			Dsn:      os.Getenv("DSN"),
			FilePath: os.Getenv("FILE_PATH"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("TOKEN"),
		},
	}
}

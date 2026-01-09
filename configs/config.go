package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db   DbConfig
	Auth AuthConfigs
}

type DbConfig struct {
	Dsn string
}

type AuthConfigs struct {
	Secret string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		Auth: AuthConfigs{
			Secret: os.Getenv("TOKEN"),
		},
	}
}

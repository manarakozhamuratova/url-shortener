package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"log"
	"time"
)

type Config struct {
	DBHost           string        `env:"DB_HOST"`
	DBPort           string        `env:"DB_PORT"`
	DBUser           string        `env:"DB_USER"`
	DBPassword       string        `env:"DB_PASS"`
	DBName           string        `env:"DB_NAME"`
	SSLMode          string        `env:"DB_SSLMODE"`
	ShortURLLen      int           `env:"SHORT_URL_LEN"`
	ShortUrlDuration time.Duration `env:"SHORT_URL_DURATION"`
}

func Load() (config Config, err error) {
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config = Config{}
	err = env.Parse(&config)
	return
}

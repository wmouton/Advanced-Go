package simplex

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Addr string
}

func LoadConfig() *Config {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		addr = ":3000"
	}

	return &Config{
		Addr: addr,
	}
}

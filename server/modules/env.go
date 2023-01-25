package modules

import (
	"github.com/joho/godotenv"
	"github.com/kataras/golog"
	"os"
)

func InitEnv() {
	golog.Info("loading .env file")
	if err := godotenv.Load(); err != nil {
		golog.Fatal("No .env file can be found.")
	}
}

func EnsureEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		golog.Fatal("Configuration ", key, " is not found, or is empty, despite being required.")
	}
	return value
}

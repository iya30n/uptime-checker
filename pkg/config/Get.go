package config

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var doOnce sync.Once

func Get(name string) string {
	doOnce.Do(func() {
		if err := godotenv.Load(".env"); err != nil {
			panic(err)
		}
	})

	return os.Getenv(name)
}

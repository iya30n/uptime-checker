package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load(".env")
    if err != nil {
        panic(err)
    }

    dbName := os.Getenv("DB_NAME")

    fmt.Println(dbName)
}
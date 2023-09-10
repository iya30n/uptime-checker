package main

import (
	"fmt"
	"uptime/pkg/config"
)

func main() {
	fmt.Println(config.Get("DB_NAME"))
}
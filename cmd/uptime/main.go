package main

import (
	"uptime/api/http"
	"uptime/internal/uptimeHandler"
)

func main() {
	uptimeHandler.Check()

	http.Serve()
}

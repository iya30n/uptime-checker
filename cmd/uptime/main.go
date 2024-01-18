package main

import (
	"uptime/api/http"
	"uptime/internal/uptimeHandler"
)

func main() {
	uptimehandler.Check()

	http.Serve()
}

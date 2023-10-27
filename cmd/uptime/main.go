package main

import (
	"uptime/api/http"
	"uptime/internal/uptimehandler"
)

func main() {
	uptimehandler.Check()

	http.Serve()
}

package influxdb

import (
	"sync"
	"uptime/pkg/config"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var doOnce sync.Once
var singletonClient influxdb2.Client

func Connect() influxdb2.Client {
	doOnce.Do(func() {
		token := config.Get("INFLUXDB_TOKEN")
		url := config.Get("INFLUXDB_URL")

		singletonClient = influxdb2.NewClient(url, token)
	})

	return singletonClient
}

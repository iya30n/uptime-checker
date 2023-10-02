package influxdb

import (
	"strings"
	"sync"
	"uptime/pkg/config"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var doOnce sync.Once
var singletonClient influxdb2.Client

func Connect() influxdb2.Client {
	doOnce.Do(func() {
		token := config.Get("INFLUXDB_TOKEN")

		singletonClient = influxdb2.NewClient(getConnectionUrl(), token)
	})

	return singletonClient
}

func getConnectionUrl() string {
	host, port := config.Get("INFLUXDB_HOST"), config.Get("INFLUXDB_PORT")

	if !strings.Contains(host, "http") {
		host = "http://" + host
	}

	return host + ":" + port
}

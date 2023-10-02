package influxdb

import (
	"context"
	"errors"
	"time"
	"uptime/pkg/config"

	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

type WriteInflux struct {
	Measurement string
	Tags  map[string]string
	Fields map[string]interface{}
}

func Write(writeOptions WriteInflux) error {
	client := Connect()

	if client == nil {
		return errors.New("influx client is nil")
	}

	writeAPI := client.WriteAPIBlocking(config.Get("INFLUXDB_ORG_NAME"), config.Get("INFLUXDB_BUCKET_NAME"))

	point := write.NewPoint(writeOptions.Measurement, writeOptions.Tags, writeOptions.Fields, time.Now())

	return writeAPI.WritePoint(context.Background(), point)
}

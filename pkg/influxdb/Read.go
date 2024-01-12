package influxdb

import (
	"context"
	"errors"
	"fmt"
	"uptime/pkg/config"
	"uptime/pkg/logger"
)

type ReadFilters struct {
	UserId    uint
	WebsiteId uint
	FromDate  string
	ToDate    string
}

func Read(filters ReadFilters) (error, []map[string]interface{}) {
	resData := []map[string]interface{}{}

	client := Connect()
	if client == nil {
		return errors.New("influx client is nil"), resData
	}

	bucket := config.Get("INFLUXDB_BUCKET_NAME")

	qApi := client.QueryAPI(config.Get("INFLUXDB_ORG_NAME"))
	query := fmt.Sprintf(`from(bucket:"%s")
		|> range(start: %s, stop: %s)
		|> filter(fn: (r) => r._measurement == "websites_uptime")
		|> filter(fn: (r) => r["_field"] == "status")
  		|> filter(fn: (r) => r["user"] == "%d")
  		|> filter(fn: (r) => r["website"] == "%d")
		|> mean()`, bucket, filters.FromDate, filters.ToDate, filters.UserId, filters.WebsiteId)

	results, err := qApi.Query(context.Background(), query)
	if err != nil {
		return err, resData
	}

	for results.Next() {
		if err := results.Err(); err != nil {
			logger.Error(err.Error())
			continue
		}

		resData = append(resData, results.Record().Values())
	}

	return nil, resData
}

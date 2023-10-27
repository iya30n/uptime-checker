package main

import (
	"encoding/json"
	"uptime/internal/jobs"
	"uptime/internal/models"
	"uptime/pkg/logger"
	"uptime/pkg/queue"
)

func main() {
	jm := models.FailedJob{}
	failedList, err := jm.Failed()
	if err != nil {
		logger.Error(err.Error())
		return
	}

	for _, j := range failedList {
		var payload jobs.JobPayload
		var err error

		err = json.Unmarshal([]byte(j.Payload), &payload)
		if err != nil {
			logger.Error(err.Error())
			break
		}

		job := jobs.Job{
			Payload: payload,
		}

		err = queue.Enqueue(j.QueueName, job)
		if err != nil {
			logger.Error(err.Error())
			return
		}

		if j.Delete() != nil {
			logger.Error(err.Error())
			return
		}
	}
}

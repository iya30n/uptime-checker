package jobs

import (
	"encoding/json"
	"uptime/internal/models"
	"uptime/pkg/logger"
)

type QueueableJob interface {
	SetData(payload JobPayload)
	Handle() bool
}

type JobPayload map[string]interface{}

type Job struct {
	Payload  JobPayload
	TryCount int
}

func (j *Job) Encode() []byte {
	bt, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}

	return bt
}

func Decode(job []byte) Job {
	var j Job
	if err := json.Unmarshal(job, &j); err != nil {
		panic(err)
	}

	return j
}

func (j *Job) Fail(queueName string) {
	ep, err := j.encodePayload()
	if err != nil {
		logger.Error(err.Error())
		return
	}

	jm := models.FailedJob{
		Status:    "failed",
		Payload:   ep,
		QueueName: queueName,
	}

	jm.Save()
}

func (j *Job) encodePayload() (string, error) {
	bt, err := json.Marshal(&j.Payload)
	if err != nil {
		return "", err
	}

	return string(bt), nil
}

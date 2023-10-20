package jobs

import "encoding/json"

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

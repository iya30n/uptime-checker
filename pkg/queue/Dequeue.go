package queue

import (
	"context"
	"uptime/internal/jobs"
	"uptime/pkg/redis"
)

func Dequeue(name string) (jobs.Job, error) {
	rc := redis.Connect()
	stringcmd := rc.RPop(context.Background(), name)

	if len(stringcmd.Val()) < 1 {
		return jobs.Job{}, stringcmd.Err()
	}

	return jobs.Decode([]byte(stringcmd.Val())), stringcmd.Err()
}

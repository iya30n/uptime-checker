package queue

import (
	"context"
	"uptime/internal/jobs"
	"uptime/pkg/redis"
)

func Enqueue(name string, job jobs.Job) error {
	rc := redis.Connect()
	intcmd := rc.LPush(context.Background(), name, job.Encode())

	return intcmd.Err()
}

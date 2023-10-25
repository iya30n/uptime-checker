package main

import (
	"context"
	"errors"
	redis2 "github.com/redis/go-redis/v9"
	"time"
	"uptime/internal/jobs"
	"uptime/pkg/logger"
	"uptime/pkg/queue"
	"uptime/pkg/redis"
)

var regJobs map[string]jobs.QueueableJob = map[string]jobs.QueueableJob{
	"upq:email": &jobs.EmailJob{},
}

var inProgress map[string]bool = make(map[string]bool)

func main() {

	rc := redis.Connect()

	for {
		scmd := rc.Keys(context.Background(), "upq:*")
		if err := scmd.Err(); err != nil {
			panic(err)
		}

		if qlen := len(scmd.Val()); qlen < 1 || qlen == len(inProgress) {
			time.Sleep(5 * time.Second)
			continue
		}

		for _, qname := range scmd.Val() {
			if _, ok := regJobs[qname]; !ok {
				continue
			}

			if _, ok := inProgress[qname]; ok {
				continue
			}

			go work(qname)
		}
	}
}

func work(name string) {
	inProgress[name] = true
	for {
		job, err := queue.Dequeue(name)

		if err != nil {
			if errors.Is(err, redis2.Nil) {
				break
			}

			logger.Error(err.Error())
			job.Failed(name)
			break
		}

		if job.Payload == nil {
			break
		}

		jt := regJobs[name]
		jt.SetData(job.Payload)
		if !jt.Handle() {
			if job.TryCount == 0 {
				job.Failed(name)
				continue
			}

			job.TryCount--

			if queue.Enqueue(name, job) != nil {
				logger.Error(err.Error())
				job.Failed(name)
				break
			}

			continue
		}
	}

	delete(inProgress, name)
}

package main

import (
	"context"
	"fmt"
	"time"
	"uptime/internal/jobs"
	"uptime/pkg/queue"
	"uptime/pkg/redis"
)

var regJobs map[string]jobs.QueueableJob = map[string]jobs.QueueableJob{
	"upq:otp": &jobs.Email{},
}

func main() {

	rc := redis.Connect()

	for {
		scmd := rc.Keys(context.Background(), "upq:*")
		if err := scmd.Err(); err != nil {
			panic(err)
		}

		if len(scmd.Val()) < 1 {
			time.Sleep(5 * time.Second)
			continue
		}

		for _, qname := range scmd.Val() {
			if _, ok := regJobs[qname]; !ok {
				continue
			}

			go work(qname)
		}
	}
}

func work(name string) {
	for {
		job, err := queue.Dequeue(name)

		if job.Payload == nil {
			continue
		}

		if err != nil {
			panic(err)

			if job.TryCount == 0 {
				// TODO: save the payload to db as failed
				fmt.Printf("job failed: %v \n", job.Encode())
				continue
			}

			job.TryCount--

			queue.Enqueue(name, job)
			continue
		}

		jt := regJobs[name]
		jt.Handle()
	}
}

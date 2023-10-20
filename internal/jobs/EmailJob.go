package jobs

import (
	"fmt"
)

type EmailJob struct {
	Data JobPayload
}

func (e *EmailJob) SetData(data JobPayload) {
	e.Data = data
}

func (e *EmailJob) Handle() bool {
	// TODO: handle should return a bool to check for retry
	fmt.Printf("sending email to: %s", e.Data["email"])
	return true
}

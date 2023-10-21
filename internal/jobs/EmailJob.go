package jobs

import (
	"uptime/pkg/logger"
	"uptime/pkg/mail"
)

type EmailJob struct {
	Data JobPayload
}

func (e *EmailJob) SetData(data JobPayload) {
	e.Data = data
}

func (e *EmailJob) Handle() bool {
	// TODO: handle should return a bool to check for retry
	email, title, view := e.Data["email"].(string), e.Data["title"].(string), e.Data["view"].(string)

	if err := mail.Send(email, title, view); err != nil {
		logger.Error(err.Error())
		return false
	}

	return true
}

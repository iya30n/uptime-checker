package auth

import (
	"fmt"
	"path/filepath"
	"uptime/internal/jobs"
	"uptime/internal/models"
	"uptime/pkg/config"
	"uptime/pkg/queue"
	"uptime/pkg/view"
)

func sendVerificationEmail(email string) error {
	otp := models.Otp{}
	code, err := otp.GenerateCode(email)
	if err != nil {
		return err
	}

	view := view.View{
		Path: filepath.Join("views", "mail", "verify.html"),
		Data: map[string]string{
			"[APP_URL]":           config.Get("APP_URL"),
			"[USER_EMAIL]":        email,
			"[VERIFICATION_CODE]": fmt.Sprintf("%d", code),
		},
	}

	job := jobs.Job{
		Payload: jobs.JobPayload{
			"email": email,
			"title": "Verification Email",
			"view":  view.Render(),
		},
	}

	return queue.Enqueue("upq:email", job)
}

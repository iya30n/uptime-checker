package mail

import (
	"net/smtp"
	"uptime/pkg/config"
)

func Send(to string, subject string, message string) error {
	message = "From: " + config.Get("SMTP_FROM") + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		message + "\r\n"

	auth := smtp.PlainAuth("", config.Get("SMTP_USERNAME"), config.Get("SMTP_PASSWORD"), config.Get("SMTP_HOST"))

	address := config.Get("SMTP_HOST") + ":" + config.Get("SMTP_PORT")

	return smtp.SendMail(address, auth, config.Get("SMTP_FROM"), []string{to}, []byte(message))
}

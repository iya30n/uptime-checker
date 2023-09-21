package mail

import (
	"net/smtp"
	"os"
	"uptime/pkg/config"
)

func Send(to string, subject string, message string) error {
	/* htmlView, err := getView("./views/mail/verify.html")
	if err != nil {
		return err
	} */

	message = "From: " + config.Get("SMTP_FROM") + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: " + "text/html; charset='utf-8'" +
		"\r\n" +
		message + "\r\n"

	auth := smtp.PlainAuth("", config.Get("SMTP_USERNAME"), config.Get("SMTP_PASSWORD"), config.Get("SMTP_HOST"))

	address := config.Get("SMTP_HOST") + ":" + config.Get("SMTP_PORT")

	return smtp.SendMail(address, auth, config.Get("SMTP_FROM"), []string{to}, []byte(message))
}

func getView(path string) (string, error) {
	content, err := os.ReadFile(path)

	return string(content), err
}

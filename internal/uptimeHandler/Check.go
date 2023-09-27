package uptimeHandler

import (
	"fmt"
	"net"
	"net/http"
	"path/filepath"
	"time"
	"uptime/internal/models"

	"uptime/pkg/influxdb"
	"uptime/pkg/logger"
	"uptime/pkg/mail"
	"uptime/pkg/view"
)

func Check() {
	websiteChan := make(chan models.Website)

	// TODO: change this to 30 minutes
	ticker := time.NewTicker(time.Minute * 3)
	go func() {
		for range ticker.C {
			// TODO: read about error handling in goroutines

			for _, w := range getWebsites() {
				websiteChan <- w
				/* status := getHttpStatus(w.Url)
				if status >= 500 {
					sendEmail(w.User.Email)
				}

				err := influxdb.Write(influxOpt(w.User, w, status))
				handleErr(err) */
			}
		}
	}()

	// worker pool
	for i := 0; i < 3; i++ {
		go worker(websiteChan)
	}
}

func worker(chn <-chan models.Website) {
	for w := range chn {
		status := getHttpStatus(w.Url)
		if status >= 500 {
			sendEmail(w.User.Email)
		}

		err := influxdb.Write(influxOpt(w.User, w, status))
		handleErr(err)
	}
}

func getWebsites() []models.Website {
	wModel := models.Website{}
	wModel.With([]string{"User"})
	websites, err := wModel.All()
	handleErr(err)

	return websites
}

func handleErr(err error) {
	if err != nil {
		logger.Error(err.Error())
		fmt.Printf("something's wrong: %s", err.Error())
	}
}

func getHttpStatus(url string) int {
	if _, err := net.ResolveIPAddr("ip", url); err != nil {
		return 500
	}

	res, err := http.Get(url)
	if err != nil {
		logger.Error(err.Error())
		return 500
	}

	defer func() {
		err := res.Body.Close()
		handleErr(err)
	}()

	return res.StatusCode
}

func influxOpt(user models.User, website models.Website, status int) influxdb.WriteInflux {
	return influxdb.WriteInflux{
		Measurement: "websites_uptime",
		Tags: map[string]string{
			"user":    fmt.Sprintf("%d", user.ID),
			"website": fmt.Sprintf("%d", website.ID),
		},
		Fields: map[string]interface{}{
			"status": status,
		},
	}
}

func sendEmail(email string) {
	view := view.View{
		Path: filepath.Join("views", "mail", "website-alert.html"),
		Data: map[string]string{
			// "[APP_URL]":           config.Get("APP_URL"),
			// "[USER_EMAIL]":        email,
			// "[VERIFICATION_CODE]": fmt.Sprintf("%d", code),
		},
	}

	err := mail.Send(email, "Website is Down!", view.Render())
	handleErr(err)
}

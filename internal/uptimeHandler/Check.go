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

	go func() {
		thirtySecondTicker := time.NewTicker(time.Second * 30)
		sixtySecondTicker := time.NewTicker(time.Minute * 1)
		fiveMinuteTicker := time.NewTicker(time.Minute * 5)
		thirtyMinuteTicker := time.NewTicker(time.Minute * 30)

		for {
			select {
			case <-thirtySecondTicker.C:
				for _, w := range getWebsites(time.Second * 30) {
					websiteChan <- w
				}

			case <-sixtySecondTicker.C:
				for _, w := range getWebsites(time.Minute * 1) {
					websiteChan <- w
				}

			case <-fiveMinuteTicker.C:
				for _, w := range getWebsites(time.Minute * 5) {
					websiteChan <- w
				}

			case <-thirtyMinuteTicker.C:
				for _, w := range getWebsites(time.Minute * 30) {
					websiteChan <- w
				}
			}
		}
	}()

	// worker pool
	for i := 0; i < 5; i++ {
		go worker(websiteChan)
	}
}

func worker(chn <-chan models.Website) {
	for w := range chn {
		status := getHttpStatus(w.Url)
		if status >= 500 && w.Notify {
			sendEmail(w.User, w.Name)
		}

		err := influxdb.Write(influxOpt(w.User, w, status))
		handleErr(err)
	}
}

func getWebsites(du time.Duration) []models.Website {
	wModel := models.Website{}
	wModel.With([]string{"User"})
	websites, err := wModel.Find("check_time = ?", du)
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

func sendEmail(user models.User, websiteName string) {
	view := view.View{
		Path: filepath.Join("views", "mail", "website-alert.html"),
		Data: map[string]string{
			"[USER]":       user.Name,
			"WEBSITE_NAME": websiteName,
		},
	}

	err := mail.Send(user.Email, "Website is Down!", view.Render())
	handleErr(err)
}

package uptimeHandler

import (
	"fmt"
	"net/http"
	"time"
	"uptime/internal/models"
	"uptime/pkg/influxdb"
)

func Check() {
	// websiteChan := make(chan string)

	ticker := time.NewTicker(time.Minute * 3)
	go func() {
		for range ticker.C {
			// TODO: read about error handling in goroutines
			wModel := models.Website{}
			wModel.With([]string{"User"})
			websites, err := wModel.All()
			if err != nil {
				panic(err)
			}
			for _, website := range websites {
				// websiteChan <- website.Url
				influxdb.Write(getInfluxOptions(website.User, website, getHttpStatus(website.Url)))
			}
		}
	}()

	// worker pool
	/* for i := 0; i < 3; i++ {
		go worker(websiteChan)
	} */
}

func getHttpStatus(url string) int {
	fmt.Printf("\n checking %s", url)

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	return res.StatusCode
}

func getInfluxOptions(user models.User, website models.Website, status int) influxdb.WriteInflux {
	// measurement := fmt.Sprintf("user_%d_website_%d", user.ID, website.ID)
	// measurement := fmt.Sprintf("user_%d_%d", user.ID, website.ID)
	// measurement := website.Url

	return influxdb.WriteInflux{
		Measurement: "websites_uptime",
		Tags: map[string]string{
			"user": fmt.Sprintf("%d", user.ID),
			"website": fmt.Sprintf("%d", website.ID),
		},
		Fields: map[string]interface{}{
			"status": status,
		},
	}
}

/* func worker(chn <-chan string) {
	for w := range chn {
		// TODO: read about error handling in goroutines
		// res, _ := http.Get(w)
	}
} */

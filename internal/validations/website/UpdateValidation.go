package website

import (
	"encoding/json"
	"fmt"
	"time"
)

type UpdateValidation struct {
	Name      string        `json:"name" binding:"required,min=3,max=100"`
	Url       string        `json:"url" binding:"required,url"`
	CheckTime time.Duration `json:"check_time" binding:"required"`
	Notify    bool          `json:"notify" binding:"required"`
}

func (cv *UpdateValidation) UnmarshalJSON(data []byte) error {
	var raw struct {
		Name      string `json:"name"`
		Url       string `json:"url"`
		CheckTime string `json:"check_time"`
		Notify    bool   `json:"notify"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	cv.Name = raw.Name
	cv.Url = raw.Url
	cv.Notify = raw.Notify

	switch raw.CheckTime {
	case "30s":
		cv.CheckTime = 30 * time.Second
	case "1m":
		cv.CheckTime = time.Minute
	case "5m":
		cv.CheckTime = 5 * time.Minute
	case "30m":
		cv.CheckTime = 30 * time.Minute
	default:
		return fmt.Errorf("check_time must be one of the following: 30s, 1m, 5m, 30m")
	}

	return nil
}
package website

import (
	"encoding/json"
	"time"
)

type CreateValidation struct {
	Name string `json:"name" binding:"required,min=3,max=100"`
	Url  string `json:"url" binding:"required,url"`
	CheckTime time.Duration `json:"check_time" binding:"required"`
}

func (cv *CreateValidation) UnmarshalJSON(data []byte) error {
    var raw struct {
        Name      string `json:"name"`
        Url       string `json:"url"`
        CheckTime string `json:"check_time"`
    }

    if err := json.Unmarshal(data, &raw); err != nil {
        return err
    }

    cv.Name = raw.Name
    cv.Url = raw.Url
    duration, err := time.ParseDuration(raw.CheckTime)
    if err != nil {
        return err
    }

    cv.CheckTime = duration
	
    return nil
}
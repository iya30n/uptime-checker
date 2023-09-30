package website

import (
	"time"
)

type UpdateValidation struct {
	Name      string `json:"name" binding:"required,min=3,max=100"`
	Url       string `json:"url" binding:"required,url"`
	CheckTime string `json:"check_time" binding:"required,oneof=30s 1m 5m 30m" time_format:"15m"`
	Notify    *bool   `json:"notify" binding:"required"`
}

func (u *UpdateValidation) GetChcekTimeDur() time.Duration {
	t, _ := time.ParseDuration(u.CheckTime)

	return t
}
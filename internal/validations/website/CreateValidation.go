package website

import "time"

type CreateValidation struct {
	Name      string `json:"name" binding:"required,min=3,max=100"`
	Url       string `json:"url" binding:"required,url"`
	CheckTime string `json:"check_time" binding:"required,oneof=30s 1m 5m 30m" time_format:"15m"`
	Notify    *bool   `json:"notify" binding:"required"`
}

func (c *CreateValidation) GetChcekTimeDur() time.Duration {
	t, _ := time.ParseDuration(c.CheckTime)

	return t
}
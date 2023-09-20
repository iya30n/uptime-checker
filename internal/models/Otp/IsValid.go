package Otp

import (
	"time"
	"uptime/pkg/database/mysql"
)

func (o Otp) IsValid(email string, code int) bool {
	db := mysql.Connect()
	currentTime := time.Now()
	res := db.Where("email = ? AND code = ? AND created_at >= ?", email, code, currentTime.Add(time.Minute-3)).Find(&o)

	if res.RowsAffected == 0 || o.Used {
		return false
	}

	return true
}

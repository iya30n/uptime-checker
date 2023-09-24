package Otp

import "uptime/pkg/database/mysql"

func (o *Otp) First(email string, code int) error {
	db := mysql.Connect()

	res := db.Where("email = ? AND code = ?", email, code).First(&o)

	return res.Error
}
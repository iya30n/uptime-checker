package Otp

import "uptime/pkg/database/mysql"

func (o *Otp) Find(email string, code int) error {
	db := mysql.Connect()

	res := db.Where("email = ? AND code = ?", email, code).Find(&o)

	return res.Error
}
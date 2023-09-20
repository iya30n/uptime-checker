package Otp

import "uptime/pkg/database/mysql"

func (o *Otp) Save() error {
	db := mysql.Connect()
	res := db.Create(&o)

	return res.Error
}

package Otp

import "uptime/pkg/database/mysql"

func (o *Otp) Update(data map[string]interface{}) error {
	db := mysql.Connect()
	res := db.Model(&o).Updates(data)
	return res.Error
}
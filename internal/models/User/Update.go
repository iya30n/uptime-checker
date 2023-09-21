package User

import "uptime/pkg/database/mysql"

func (u *User) Update(data map[string]interface{}) error {
	db := mysql.Connect()
	res := db.Model(&u).Updates(data)

	return res.Error
}
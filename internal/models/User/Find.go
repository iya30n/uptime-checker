package User

import "uptime/pkg/database/mysql"

func (u *User) Find(key string, value any) error {
	db := mysql.Connect()

	res := db.Where("username = ? OR email = ?", u.Username, u.Email).Find(&u)

	return res.Error
}
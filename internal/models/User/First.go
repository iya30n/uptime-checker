package User

import "uptime/pkg/database/mysql"

func (u *User) First(query string, values ...interface{}) error {
	db := mysql.Connect()

	res := db.Where(query, values...).First(&u)

	return res.Error
}

package User

import "uptime/pkg/database/mysql"

func (u *User) Find(query string, values ...interface{}) error {
	db := mysql.Connect()

	res := db.Where(query, values...).Find(&u)

	return res.Error
}
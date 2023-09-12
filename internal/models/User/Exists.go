package User

import "uptime/pkg/database/mysql"

func (u *User) Exists() (bool, error) {
	db := mysql.Connect()

	res := db.Where("username = ? OR email = ?", u.Username, u.Email).Find(&User{})

	return res.RowsAffected != 0, res.Error
}

package User

import (
	"uptime/pkg/database/mysql"

	"golang.org/x/crypto/bcrypt"
)

func (u *User) Save() error {
	passwd, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	u.Password = string(passwd)

	db := mysql.Connect()
	res := db.Create(&u)

	return res.Error
}

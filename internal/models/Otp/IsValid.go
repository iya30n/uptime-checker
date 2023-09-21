package Otp

import "time"

// "time"
// "uptime/pkg/database/mysql"

func (o Otp) IsValid() bool {
	/* db := mysql.Connect()
	// currentTime := time.Now()
	// res := db.Where("email = ? AND code = ? AND created_at >= ?", email, code, currentTime.Add(time.Minute-3)).Find(&o)
	res := db.Where("email = ? AND code = ?", email, code).Find(&o)

	if res.RowsAffected == 0 || o.Used {
		return false
	}

	return true */

	return !o.Used && o.CreatedAt.After(time.Now().Add(-time.Minute * 3))
}

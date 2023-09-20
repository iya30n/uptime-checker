package main

import (
	"fmt"
	"uptime/internal/models/Otp"
	"uptime/internal/models/User"
	"uptime/internal/models/Website"
	"uptime/pkg/database/mysql"
)

func main() {
	migrations := map[string]interface{}{
		"user":              User.User{},
		"website":           Website.Website{},
		"Otp": Otp.Otp{},
	}

	db := mysql.Connect()

	for name, migration := range migrations {
		fmt.Printf("migrating %s \n", name)
		db.AutoMigrate(migration)
	}

	fmt.Println("Done!")
}

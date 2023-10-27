package main

import (
	"fmt"
	"uptime/internal/models"
	"uptime/pkg/database/mysql"
)

func main() {
	db := mysql.Connect()

	models := []interface{}{
		models.User{},
		models.Website{},
		models.Otp{},
		models.FailedJob{},
	}

	if err := db.AutoMigrate(models...); err != nil {
		panic(err)
	}

	fmt.Println("Done!")
}

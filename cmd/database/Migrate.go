package database

import (
	"fmt"
	"uptime/internal/models/User"
	"uptime/pkg/database/mysql"
)

func Migrate() {
	migrations := map[string]interface{}{
		"user": User.User{},
	}

	db := mysql.Connect()

	for name, migration := range migrations {
		fmt.Printf("migrating %s \n", name)
		db.AutoMigrate(migration)
	}

	fmt.Println("Done!")
}

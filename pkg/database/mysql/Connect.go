package mysql

import (
	"fmt"
	"sync"
	"uptime/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var doOnce sync.Once
var singletonConnection *gorm.DB

func Connect() *gorm.DB {
	doOnce.Do(func() {
		dsn := fmt.Sprintf(
			"%s:@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Get("DB_USERNAME"),
			config.Get("DB_HOST"),
			config.Get("DB_PORT"),
			config.Get("DB_NAME"),
		)

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			SkipDefaultTransaction: true,
		})
		if err != nil {
			panic(err)
		}

        singletonConnection = db
	})

    return singletonConnection
}

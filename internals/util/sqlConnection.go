package util

import (
	"fmt"
	"os"
	"sync"
	"time"
	"truthly/internals/util/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	once sync.Once
	db   *gorm.DB
)

func InitDb() *gorm.DB {
	once.Do(func() {
		log := logger.InitLogger()

		// Get the variables from env
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")

		// now make dsn
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			user, password, host, port, dbname,
		)

		var err error

		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Error("Error in connecting with the mysql: " + err.Error())
			panic(err.Error())
		}

		log.Info("Mysql connection seccessful")

		sqlDb, _ := db.DB()
		sqlDb.SetMaxOpenConns(20)
		sqlDb.SetConnMaxIdleTime(10)
		sqlDb.SetConnMaxLifetime(time.Hour)
	})

	return db
}

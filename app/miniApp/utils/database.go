package utils

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"sync"

	"gorm.io/gorm"
)

var (
	dsn string = "root:123456@tcp(127.0.0.1:3306)/miniapp?charset=utf8mb4&parseTime=True&loc=Local"
)
var db *gorm.DB

// 调用一次
var once sync.Once

func GetDb() *gorm.DB {
	once.Do(func() {
		dba, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			logrus.Fatal(err)
		}
		db = dba
		//dba, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
		//if err != nil {
		//	logrus.Fatal(err)
		//}
		//db = dba
	})

	return db
}

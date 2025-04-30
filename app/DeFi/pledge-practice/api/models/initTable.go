package models

import "p-cy/db"

func InitTable() {
	db.Mysql.AutoMigrate(&User{})
	db.Mysql.AutoMigrate(&TEmail{})
}

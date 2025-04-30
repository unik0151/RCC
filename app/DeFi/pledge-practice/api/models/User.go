package models

import (
	"errors"
	"p-cy/db"
	"p-cy/log"
)

type User struct {
	Id       uint64     `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Key      string     `gorm:"column:key;unique" json:"key" `
	UserName string     `gorm:"column:user_name" json:"name" `
	Password string     `gorm:"column:password" json:"password" `
	Type     UserStatus `gorm:"column:type" json:"type" `
}

type UserStatus int

var (
	COMMON_USER_STATUS_UNKNOWN UserStatus = 1
	ADMIN                      UserStatus = 999
)

func NewUser() *User {
	return &User{}
}

//	func (User) TableName() string  {
//		return "user"
//	}
func (u *User) FindInfoByName(name string) (*User, error) {
	user := &User{}
	err := db.Mysql.Table("user").Where("user_name = ?", name).First(user).Debug().Error
	if err != nil {
		log.Logger.Sugar().Error("findInfoByKey", "name", name, "err", err)
		return nil, err
	}
	return user, nil
}

func (e *User) FindEmail(email string) (*User, error) {
	var result User
	err := db.Mysql.Table("user").Where("email = ?", email).First(&result).Debug().Error
	if err != nil {
		return nil, errors.New("user not exists")
	}
	return &result, nil
}

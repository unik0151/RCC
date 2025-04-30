package models

import (
	"errors"
	"p-cy/api/common/statuscode"
	"p-cy/config"
	"p-cy/db"
	"time"
)

type TEmail struct {
	Code       string        `gorm:"primary_key;column:email_code"`
	CreateDate time.Time     `gorm:"column:create_date"`
	Email      string        `gorm:"column:email"`
	Duration   time.Duration `gorm:"column:duration"`
	EmailType  int           `gorm:"column:email_type"`
}

func (TEmail) TableName() string {
	return "t_email"
}

func NewEmail() *TEmail {
	return &TEmail{}
}

func (e *TEmail) BeforeSave() {
	e.CreateDate = time.Now()
}

//func (e *TEmail) AfterSave() {
//
//}

func (e *TEmail) SaveEmail(email, code string) int {

	//发送code
	e.Email = email
	e.Code = code
	e.CreateDate = time.Now()
	e.Duration = time.Duration(config.GlobalConfig.Email.ValidateTime) * time.Minute
	err := db.Mysql.Table("t_email").Save(e).Debug().Error
	if err != nil {
		return statuscode.ServerError
	}
	return statuscode.CommonSuccess
}

func (u *TEmail) FindInfoByEmail(email string) (*TEmail, error) {
	var emailE TEmail
	err := db.Mysql.Table("t_email").Where("email = ?", email).Order("create_date desc").First(&emailE).Debug().Error
	if err != nil {
		return nil, errors.New("email not found")
	}
	return &emailE, nil
}

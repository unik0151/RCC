package model

import (
	"github.com/sirupsen/logrus"
	"miniapp/utils"
)

type AccountB struct {
	Path     string `gorm:"column:path;type:varchar(255)" json:"path"`
	Password string `gorm:"column:password;type:varchar(255)" json:"password"`
	Address  string `gorm:"column:address;type:varchar(255)" json:"address"`
	BaseModel
}

func NewAccountB(path string, password string, address string) *AccountB {
	return &AccountB{
		Path:     path,
		Password: password,
		Address:  address,
	}
}

func (*AccountB) TableName() string {
	return "account"
}

func init() {
	db := utils.GetDb()

	err := db.AutoMigrate(&AccountB{})
	if err != nil {
		logrus.Error(err)
	}

}

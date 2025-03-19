package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义实体
type FileDTO struct {
	Creator  string `gorm:"column:creator;type:varchar;size:255"`
	FilePath string `gorm:"column:file_path;type:varchar;size:255"`
	gorm.Model
}

func init() {
	GetDBDriver().AutoMigrate(FileDTO{})
}

// 选择使用哪个数据库，这里用sqllit

func GetDBDriver() *gorm.DB {

	dsn := "root:123456@tcp(127.0.0.1:3306)/go_database?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		CreateBatchSize: 1000,
	})

	if err != nil {
		panic("Database connect Error !")

	}

	return db
}

func CreateDATA(gorm *gorm.DB, file *FileDTO) (db *gorm.DB) {

	var temp []FileDTO
	gorm.Find(&temp)
	for _, v := range temp {
		fmt.Printf("v: %v\n", v)
	}

	return gorm.Create(file)
}

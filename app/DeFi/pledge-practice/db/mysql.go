package db

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"p-cy/config"
	"p-cy/log"
	"time"
)

func InitMysql() {
	globalConfig := config.GlobalConfig

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       globalConfig.MySql.Dsn,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportDropConstraint: true,
		DontSupportRenameColumn:   true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Logger.Error("mysql init fail", zap.Error(err))
	}

	_ = db.Callback().Create().After("gorm:after_create").Register("after_create", After)
	_ = db.Callback().Update().After("gorm:after_update").Register("after_update", After)
	_ = db.Callback().Delete().After("gorm:after_delete").Register("after_delete", After)
	_ = db.Callback().Query().After("gorm:after_query").Register("after_query", After)

	sqlDb, err := db.DB()
	if err != nil {
		log.Logger.Error("gorm mysql db init fail", zap.Error(err))
	}
	sqlDb.SetMaxIdleConns(globalConfig.MySql.MaxIdleConns)
	sqlDb.SetConnMaxIdleTime(time.Duration(globalConfig.MySql.MaxLifeTime) * time.Second)
	sqlDb.SetMaxOpenConns(globalConfig.MySql.MaxOpenConns)
	Mysql = db
}

/*
*
输出
*/
func After(db *gorm.DB) {
	db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
}

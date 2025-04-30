package main

import (
	"github.com/gin-gonic/gin"
	"p-cy/api/models"
	"p-cy/api/router"
	"p-cy/api/validate"
	"p-cy/config"
	"p-cy/db"
	"strconv"
)

func main() {
	db.InitMysql()

	models.InitTable()

	//添加校验器
	validate.BindValidate()

	engine := gin.Default()
	//中间件先不加
	router.InitRouter(engine)

	engine.Run(":" + strconv.Itoa(config.GlobalConfig.Env.Port))
}

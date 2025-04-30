package main

import (
	"miniapp/conf"
	"miniapp/middleware"
	_ "miniapp/model"
	"miniapp/router"

	"github.com/gin-gonic/gin"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	//添加可跨域请求，日志打印
	engine.Use(middleware.Next(), middleware.LoggerTOFile(), middleware.Recover())
	router.InitRouter(engine)
	engine.Run(conf.PORT)
}

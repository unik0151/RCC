package router

import (
	"github.com/gin-gonic/gin"
	"p-cy/api/controller"
	"p-cy/api/middleware"
	"p-cy/config"
	"strconv"
)

func InitRouter(e *gin.Engine) {
	v2Group := e.Group("/api/v" + strconv.Itoa(config.GlobalConfig.Env.Version))

	userController := controller.UserController{}
	v2Group.POST("/user/login", userController.Login)

	v2Group.POST(
		"/user/update",
		middleware.JwtCheck(),
		userController.Update,
	)
	emailController := controller.EmailController{}
	//code
	v2Group.POST(
		"/code",
		emailController.EmailSend,
	)

}

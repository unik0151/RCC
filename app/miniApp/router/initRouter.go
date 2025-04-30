package router

import (
	"github.com/gin-gonic/gin"
	"miniapp/api"
	"miniapp/middleware"
)

func InitRouter(r *gin.Engine) {

	auth := r.Group("/user")
	auth.POST("/login", api.Login)

	account := r.Group("/account")
	account.Use(middleware.AuthMiddleware())
	account.GET("/get", api.Select)
	account.POST("/createAccount", api.CreateAccount)
	account.GET("/get/:address", api.QueryAddressTransaction)
	account.POST("/transfer", api.Transfer)

}

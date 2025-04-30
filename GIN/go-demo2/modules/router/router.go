package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 实现一个 父接口定义，子接口定义，方法 统一放入 ， single 标识单独接口
func InitRouter(r *gin.Engine) {

	v11 := r.Group("/v1")

	v11.GET("/user").Use(func(ctx *gin.Context) {
		URI := ctx.Request.RequestURI
		body := ctx.Request.Body
		client := ctx.ClientIP()

		fmt.Printf("URI: %v\n", URI)
		fmt.Printf("body: %v\n", body)
		fmt.Printf("client: %v\n", client)

	})

}

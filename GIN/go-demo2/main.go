package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	engin := gin.Default()
	// 声明中间件层
	engin.Use()

}

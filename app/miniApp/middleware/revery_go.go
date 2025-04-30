package middleware

import (
	"fmt"
	"miniapp/utils"

	"github.com/gin-gonic/gin"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				utils.Panic(fmt.Sprintf("%s", err))
			}
		}()
	}
}

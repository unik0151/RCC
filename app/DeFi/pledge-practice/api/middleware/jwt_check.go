package middleware

import (
	"github.com/gin-gonic/gin"
	"p-cy/api/common/statuscode"
	"p-cy/api/models/response"
	"p-cy/util"
)

func JwtCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := response.Gin{Ctx: c}
		token := c.GetHeader("authorization")

		username, err := util.ParseToken(token)
		if err != nil {
			res.Response(statuscode.TokenError, nil)
			c.Abort()
			return
		}
		c.Set("username", username)

		c.Next()
	}
}

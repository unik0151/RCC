package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"miniapp/model"
	"miniapp/service"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.ResultData{
				Code:    403,
				Message: "未认证",
			})
			return
		}
		claim := &model.Claim{}
		token, err := jwt.ParseWithClaims(tokenStr, claim, func(token *jwt.Token) (interface{}, error) {
			return service.GetJwtKey(), nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.ResultData{})
		}

		c.Set("claim", claim)
		c.Next()

	}
}

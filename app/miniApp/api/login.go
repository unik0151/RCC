package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"miniapp/model"
	"miniapp/service"
	"net/http"
)

func Login(c *gin.Context) {
	var user model.Login
	err := c.ShouldBind(&user)
	if err != nil {
		logrus.Error(err)
	}
	if user.Username == "" || user.Password == "" {
		c.JSON(http.StatusUnauthorized, model.ResultData{
			Code:    500,
			Message: "用户名或密码不能为空",
		})
	}
	if u := model.Users[user.Username]; u == "" {
		c.JSON(http.StatusOK, model.ResultData{
			Code:    400,
			Message: fmt.Sprintf("该用户:【%s】不存在", user.Username),
		})
		return
	}
	token, err := service.GenerateToken(user.Username)
	if err != nil {
		logrus.Error(err)
	}
	c.JSON(http.StatusOK, model.ResultData{
		Code:    200,
		Message: "success",
		Data:    token,
	})

}

package controller

import (
	"github.com/gin-gonic/gin"
	"p-cy/api/common/statuscode"
	"p-cy/api/models/response"
	"p-cy/api/service"
	"p-cy/api/validate"
)

type EmailController struct{}

func NewEmailController() *EmailController {
	return &EmailController{}
}

func (e EmailController) EmailSend(gin *gin.Context) {

	resp := response.Gin{Ctx: gin}
	param := gin.Param("email")

	//1. 校验邮箱
	errCode := validate.NewEmail().Validate(param)
	if statuscode.CommonSuccess != errCode {
		resp.Response(errCode, nil)
		return
	}

	//2.发送逻辑
	errCode = service.NewEmail().Send(param)
	if statuscode.CommonSuccess != errCode {
		resp.Response(errCode, nil)
		return
	}
}

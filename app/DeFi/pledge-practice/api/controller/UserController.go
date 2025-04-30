package controller

import (
	"github.com/gin-gonic/gin"
	"p-cy/api/common/statuscode"
	"p-cy/api/models/request"
	"p-cy/api/models/response"
	"p-cy/api/service"
	"p-cy/api/validate"
)

type UserController struct{}

func NewUserController() *UserController {
	return new(UserController)
}

func (UserController *UserController) Login(ctx *gin.Context) {
	r := response.Gin{Ctx: ctx}
	user := request.User{}

	resp := response.User{}
	code := validate.CheckUser(ctx, &user)

	if statuscode.CommonSuccess != code {
		r.Response(code, nil)
		return
	}
	errorCode := service.NewUserService().Login(&user, &resp)
	if statuscode.CommonSuccess != errorCode {
		r.Response(errorCode, nil)
		return
	}

	r.Response(errorCode, resp)

}

func (UserController *UserController) Update(ctx *gin.Context) {
	r := response.Gin{Ctx: ctx}
	user := request.UpdateUser{}

	//校验参数
	validateCode := validate.CheckUpdateUser(ctx, &user)
	if statuscode.CommonSuccess != validateCode {
		r.Response(validateCode, nil)
		return
	}
	//Service对比
	updateCode := service.NewUserService().Update(&user, &r)
	
	r.Response(updateCode, nil)
}

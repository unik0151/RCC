package validate

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"p-cy/api/common/statuscode"
	"p-cy/api/models/request"
	"strings"
)

func CheckUser(ctx *gin.Context, user *request.User) int {

	err := ctx.ShouldBind(&user)
	if err == io.EOF {
		return statuscode.ParamError
	} else if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			if e.Field() == "Username" && e.Tag() == "required" {
				return statuscode.ParamError
			}
			if e.Field() == "Password" && e.Tag() == "required" {
				return statuscode.ParamError
			}
		}
		return statuscode.ServerError
	}
	return statuscode.CommonSuccess
}

func CheckUpdateUser(ctx *gin.Context, r *request.UpdateUser) int {
	err := ctx.ShouldBind(r)
	if err == io.EOF {
		return statuscode.ParamError
	}
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			if e.Field() == "Password" && strings.Contains(e.Tag(), "required") {
				return statuscode.ParamError
			}
			if e.Field() == "RePassword" && strings.Contains(e.Tag(), "required") {
				return statuscode.ParamError
			}
			if e.Field() == "Code" && strings.Contains(e.Tag(), "required") {
				return statuscode.ParamError
			}
		}
		return statuscode.ServerError
	}
	return statuscode.CommonSuccess

}

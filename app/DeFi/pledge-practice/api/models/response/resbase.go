package response

import (
	"github.com/gin-gonic/gin"
	"p-cy/api/common/statuscode"
)

type Response struct {
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type Gin struct {
	Ctx *gin.Context
}

func (gin *Gin) Response(code int, data interface{}, status ...int) {
	lang := statuscode.Lang_en

	value, exists := gin.Ctx.Get("lang")

	if exists {
		lang = value.(int)
	}
	res := Response{
		Code: code,
		Msg:  statuscode.GetMsg(code, lang),
		Data: data,
	}
	HttpStatus := 200
	if len(status) > 0 {
		HttpStatus = status[0]
	}
	gin.Ctx.JSON(HttpStatus, res)

}

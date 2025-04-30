package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"miniapp/model"
	"miniapp/service"
	"net/http"
)

func Select(ctx *gin.Context) {
	data := service.SelectAll()
	anies := data.Data.([]model.AccountB)
	for i, a := range anies {
		fmt.Println(i, a)
	}
	ctx.JSON(http.StatusOK, data)
}

func CreateAccount(ctx *gin.Context) {
	var re map[string]interface{}
	err := ctx.ShouldBindJSON(&re)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pass := re["password"]
	account, secret := service.CreateAccount(pass.(string))
	data := map[string]interface{}{
		"walletAddress": account.Address.Hex(),
		"secret":        secret,
	}
	ctx.JSON(http.StatusOK, model.ResultData{Code: 200, Data: data, Message: "请求成功"})
}

func QueryAddressTransaction(context *gin.Context) {
	//param := context.Param("address")

}

func Transfer(context *gin.Context) {

}

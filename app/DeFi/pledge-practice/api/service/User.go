package service

import (
	"p-cy/api/common/statuscode"
	"p-cy/api/models"
	"p-cy/api/models/request"
	"p-cy/api/models/response"
	"p-cy/log"
	"p-cy/util"
	"time"
)

type UserService struct{}

func NewUserService() *UserService {
	return new(UserService)
}

func (userService *UserService) Login(user *request.User, re *response.User) int {
	u, err := models.NewUser().FindInfoByName(user.Username)

	if err != nil || user.Password != u.Password {
		return statuscode.PasswordOrUsernameError
	}
	//创建token
	token, err := util.CreateToken(user.Username)
	if err != nil {
		log.Logger.Error("CreateToken faild !" + err.Error())
		return statuscode.ServerError
	}
	re.TokenId = token

	return statuscode.CommonSuccess
}

func (userService *UserService) Update(req *request.UpdateUser, resp *response.Gin) int {
	email := req.Email
	_, err := models.NewUser().FindEmail(email)
	if err != nil {
		log.Logger.Error("FindEmail faild !" + err.Error())
		return statuscode.UserExistError
	}
	//1. 用户密码和传入对比
	if req.Password != req.RePassword {
		return statuscode.ParamError
	}

	tEmail, err := models.NewEmail().FindInfoByEmail(email)
	if err != nil {
		log.Logger.Error("FindEmail faild !" + err.Error())
	}
	if tEmail.Code != req.Code {
		return statuscode.CodeError
	}
	if tEmail.CreateDate.Add(tEmail.Duration).Before(time.Now()) {
		return statuscode.CodeError
	}
	return statuscode.CommonSuccess
}

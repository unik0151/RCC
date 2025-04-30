package service

import (
	"p-cy/api/common/statuscode"
	"p-cy/api/models"
	"p-cy/log"
	"p-cy/util"
	"time"
)

type Email struct{}

func NewEmail() *Email {
	return &Email{}
}

/*
*
1. 找寻系统有无此邮箱
2. 有无发送验证码
3. 生成code
4. 保存系统
*/
func (email *Email) Send(emal string) int {
	var emailE models.TEmail
	var code = util.GenerateCode(6)
	//1. 系统有此邮箱 这块放个缓存redis
	_, errUser := models.NewUser().FindEmail(emal)
	if errUser != nil {
		return statuscode.UserExistError
	}
	//2. 没法送过，接着发 ， 发过，且在五分钟内，回复验证码已发送，请五分钟内输入确认
	_, errEmail := models.NewEmail().FindInfoByEmail(emal)
	//说明不存在直接发送
	if errEmail != nil {
		//发送验证码
		codeSend(code)
		//入库
		models.NewEmail().SaveEmail(emal, code)
		return statuscode.CommonSuccess
	}
	if emailE.CreateDate.Add(emailE.Duration).Before(time.Now()) {
		return statuscode.CodeExistError
	}
	codeSend(code)
	//入库
	models.NewEmail().SaveEmail(emal, code)
	return statuscode.CommonSuccess
}

func codeSend(code string) {
	go func() {
		err := util.SendEmail([]byte(code), 1)
		if err != nil {
			log.Logger.Error("send email fail")
		}
	}()
}

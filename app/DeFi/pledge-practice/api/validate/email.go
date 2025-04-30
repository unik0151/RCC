package validate

import (
	"p-cy/api/common/statuscode"
	"p-cy/util"
)

type Email struct{}

func NewEmail() *Email {
	return &Email{}
}

func (e *Email) Validate(email string) int {
	ok := util.CheckEmailRegular(email)
	if ok {
		return statuscode.CommonSuccess
	}
	return statuscode.ParamError
}

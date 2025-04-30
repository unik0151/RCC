package model

type ResultData struct {
	Code    Code        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Code int

const (
	Success Code = 200
	Error   Code = 400
)

// 给对象本身添加方法
func (re *ResultData) SetCode(code Code) *ResultData {
	re.Code = code
	return re
}

func (re *ResultData) SetMessage(message string) *ResultData {
	re.Message = message
	return re
}

func (re *ResultData) SetData(data interface{}) *ResultData {
	re.Data = data
	return re
}

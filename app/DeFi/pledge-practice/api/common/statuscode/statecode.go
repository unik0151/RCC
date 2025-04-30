package statuscode

var (
	Lang_en    = 1000
	Lang_zh    = 1001
	Lang_zh_tw = 1002

	// Login 2000 - 2100
	LoginError              = 2000
	PasswordOrUsernameError = 2000

	//param 2100 -2200
	ParamError     = 2101
	TokenError     = 2102
	NullError      = 2103
	UserExistError = 2104

	CodeError      = 2105
	CodeExistError = 2106
	// 9000-10000
	CommonSuccess = 9000
	ServerError   = 9001
)

var Msg = map[int]map[int]string{
	CommonSuccess: {
		Lang_zh: "成功!",
		Lang_en: "SUCCESS",
	},
	NullError: {
		Lang_zh: "不能为空!",
		Lang_en: "not null",
	},
	CodeError: {
		Lang_zh: "验证码无效!",
		Lang_en: "code invalid",
	},
	CodeExistError: {
		Lang_zh: "验证码发送太频繁，请五分钟内输入验证码!",
		Lang_en: "code already exist , please input code",
	},
	UserExistError: {
		Lang_zh: "邮箱用户不存在!",
		Lang_en: "email mapping user not exists",
	},
	TokenError: {
		Lang_zh: "权限不足！",
		Lang_en: "no permission",
	},
	LoginError: {
		Lang_zh:    "登陆失败",
		Lang_en:    "login failed",
		Lang_zh_tw: "操作失败",
	},
	ServerError: {
		Lang_en: "server access error , please try again later",
		Lang_zh: "服务器错误 ，请稍后重试",
	},
	ParamError: {
		Lang_zh: "参数错误",
		Lang_en: "param invalidate",
	},
	PasswordOrUsernameError: {
		Lang_zh: "用户名或密码错误",
		Lang_en: "username or password error",
	},
}

func GetMsg(code, lang int) string {
	_, ok := Msg[code]
	if ok {
		v, err := Msg[code][lang]
		if !err {
			return Msg[ServerError][lang]
		}
		return v
	}
	return Msg[ServerError][lang]
}

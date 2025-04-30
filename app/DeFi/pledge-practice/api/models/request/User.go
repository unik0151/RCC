package request

type User struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type UpdateUser struct {
	Email      string `json:"email" form:"email" binding:"required;IsEmail"`
	Password   string `json:"password" form:"password" binding:"required;IsPassword"`
	RePassword string `json:"re_password" form:"re_password" binding:"required;IsPassword"`
	Code       string `json:"code" form:"code" binding:"required"`
}

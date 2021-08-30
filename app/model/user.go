package model

// 用户模型
type User struct {
	Id       int64  `form:"id" json:"id" validate:"numeric" swaggerignore:"true"`                      // ID
	UserName string `form:"username" json:"username" uri:"username" xml:"username" binding:"required"` // 用户名
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"` // 密码
}

// 用户登录
func (u *User) Login() (user_id int64, msg string) {
	if u.UserName != "test" || u.Password != "888" {
		return 0, "用户名或密码错误"
	}
	return 1, "登录成功"
}

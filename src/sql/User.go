package sql

// User is user object
type User struct {
	// 数据ID号
	id int64

	// 用户名
	name string

	// 年龄
	age int

	// 性别
	sex int

	// 用户名
	UserName string

	// 密码
	Password string

	// 图片
	icon string

	// 个性签名
	signature string

	// 创建日期
	createDate string

	// 最后一次登录日期
	lastDate string
}

// NewUser is new user
// params:
// 	username(string) is user name
// 	password(string) is pass word
func NewUser(username, password string) *User {
	var user = new(User)
	user.UserName = username
	user.Password = password
	return user
}

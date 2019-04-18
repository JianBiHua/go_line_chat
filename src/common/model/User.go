package model

// User is user object
type User struct {
	// 数据ID号, 不导出到JSON
	ID int64 `json:"-"`

	// 用户名
	Name string `json:"name"`

	// 年龄
	Age string `json:"age"`

	// 性别
	Sex int `json:"sex"`

	// 用户名
	UserName string `json:"user"`

	// 密码, 不导出到JSON
	Password string `json:"-"`

	// 图片
	Icon string `json:"icon"`

	// 个性签名
	Signature string `json:"signature"`

	// 创建日期
	CreateDate string `json:"createDate"`

	// 最后一次登录日期
	LastDate string `json:"lastDate"`

	// 状态， 不导出到json
	Status int `json:"-"`

	// 进组日期
	JoinDate string `json:"joinDate"`
}

// NewUser is new user
// params:
// 	username(string) is user name
// 	password(string) is pass word
// out:
//  a user object pointer
func NewUser(username, password string) *User {
	var user = new(User)
	user.UserName = username
	user.Password = password
	return user
}

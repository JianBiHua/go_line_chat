package model

// UserChatGroup is user chat group object
type UserChatGroup struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Signature  string `json:"signature"`
	Max        int64  `json:"max"`
	GroupType  int64  `json:"groupType"`
	Icon       string `json:"icon"`
	CreateDate string `json:"createDate"`

	// 成员列表
	Member []*User `json:"users"`
}

package model

// UserChatMessage is user chat message object
type UserChatMessage struct {
	Id       int64  `json:"id"`
	UserName string `json:"user"`
	SendDate string `json:"sendDate"`
	Comment  string `json:"msg"`
	MsgType  int64  `json:"msgType"`
}

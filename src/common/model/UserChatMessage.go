package model

// UserChatMessage is user chat message object
type UserChatMessage struct {
	ID       int64  `json:"id"`
	UserName string `json:"user"`
	GroupID  int64  `json:"group"`
	SendDate string `json:"sendDate"`
	Comment  string `json:"msg"`
	MsgType  int64  `json:"msgType"`
}

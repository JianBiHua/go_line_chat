package event

import (
	"net"
)

// EventFriendDel is sign in event
type EventFriendDel struct {
	client net.Conn
}

// 请求消息:
// {
//	"type" : "Del",
//	"user" : "user1"
// }

// Parse is parse del friend json string, and return result string
func (a *EventFriendDel) Parse(jsonReslut map[string]interface{}, client net.Conn) {
	//
	a.client = client

}

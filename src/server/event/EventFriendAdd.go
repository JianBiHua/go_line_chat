package event

import (
	"net"
)

// EventFriendAdd is sign in event
type EventFriendAdd struct {
	client net.Conn
}

// 请求消息:
// {
//	"type" : "Add",
//	"user" : "user1"
// }

// Parse is parse add friends json string, and return result string
func (a *EventFriendAdd) Parse(jsonReslut map[string]interface{}, client net.Conn) {
	//
	a.client = client

	//
}

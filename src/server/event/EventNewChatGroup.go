package event

import (
	"net"
)

// EventNewChatGroup is new Chat Group event
type EventNewChatGroup struct {
	client net.Conn
}

// Parse is parse new Chat group json string, and return result string
func (a *EventNewChatGroup) Parse(jsonReslut map[string]interface{}, client net.Conn) {
	//
	a.client = client

}

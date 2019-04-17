package event

import (
	"net"
)

// EventChatGroupCreate is new Chat Group event
type EventChatGroupCreate struct {
	client net.Conn
}

// Parse is parse create Chat group json string, and return result string
func (a *EventChatGroupCreate) Parse(jsonReslut map[string]interface{}, client net.Conn) {
	//
	a.client = client

	//
	
	
}

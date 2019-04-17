package event

import (
	"net"
)

// EventMsg is sign in event
type EventMsg struct {
	client net.Conn
}

// Parse is parse msg json string, and return result string
func (a *EventMsg) Parse(jsonReslut map[string]interface{}, client net.Conn) {
	//
	a.client = client

	//
}

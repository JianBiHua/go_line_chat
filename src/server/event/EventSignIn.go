package event

import "net"

// EventSignIn is sign in event
type EventSignIn struct {
	client net.Conn
}

// Parse is parse sign in json string, and return result string
func (a *EventSignIn) Parse(jsonReslut map[string]interface{}, client net.Conn) {
	//
	a.client = client
}

package event

import (
	"net"
)

// EventChat is chat event
type EventChat struct {
	client net.Conn
}

// Parse is parse chat json string, and return result string
func (a *EventChat) Parse(jsonReslut map[string]interface{}, client net.Conn) {
	//
	a.client = client

}

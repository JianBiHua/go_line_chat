package event

import "net"

// EventInterface is event interface
type EventInterface interface {
	// Parse is parse atom and send msg
	Parse(jsonReslut map[string]interface{}, client net.Conn)
}

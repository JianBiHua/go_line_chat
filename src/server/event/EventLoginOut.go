package event

import (
	"net"
)

// EventLoginOut is login out event
type EventLoginOut struct {
	client net.Conn
}

// Parse is parse login out json string, and return result string
func (a *EventLoginOut) Parse(jsonReslut map[string]interface{}, client net.Conn) {
	//
	a.client = client

	// 退出成功
	client.Write([]byte("{\"type\":\"LoginOut\",\"result\":\"OK\"}\r\n"))
}

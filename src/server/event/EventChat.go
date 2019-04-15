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

	// cmd := fmt.Sprintf("INSERT INTO %s(userName, password) values(\"%v\",\"%v\")",
	// 	sql.SQLTableUser, jsonReslut["user"], jsonReslut["psw"])
	// _, err := sql.GetInstance().Insert(cmd)
	// if err != nil {
	// 	client.Write([]byte("{\"type\":\"SignIn\",\"result\":\"Error\"}\r\n"))
	// } else {
	// 	client.Write([]byte("{\"type\":\"SignIn\",\"result\":\"OK\"}\r\n"))
	// }
}

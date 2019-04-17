package event

import (
	"GoWorkspace/go_line_chat/src/server/sql"
	"fmt"
	"net"
)

// EventSignIn is sign in event
type EventSignIn struct {
	client net.Conn
}

// Parse is parse sign in json string, and return result string
func (a *EventSignIn) Parse(jsonReslut map[string]interface{}, client net.Conn) {
	//
	a.client = client

	// 插入到数据库
	cmd := fmt.Sprintf("INSERT INTO %s(userName, password, name) values(\"%v\",\"%v\",\"%v\")",
		sql.SQLTableUser, jsonReslut["user"], jsonReslut["psw"], jsonReslut["user"])
	_, err := sql.GetInstance().Insert(cmd)
	if err != nil {
		client.Write([]byte("{\"type\":\"SignIn\",\"result\":\"user sigin in error!\"}\r\n"))
	} else {
		client.Write([]byte("{\"type\":\"SignIn\",\"result\":\"OK\"}\r\n"))
	}
}

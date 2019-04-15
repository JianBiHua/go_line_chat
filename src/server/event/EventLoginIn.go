package event

import (
	"GoWorkspace/go_line_chat/src/server/sql"
	"fmt"
	"net"
)

// EventLoginIn is sign in event
type EventLoginIn struct {
	client net.Conn
}

// Parse is parse login in json string, and return result string
func (a *EventLoginIn) Parse(jsonReslut map[string]interface{}, client net.Conn) {
	//
	a.client = client

	if jsonReslut["user"] == nil || jsonReslut["psw"] == nil {
		client.Write([]byte("{\"type\":\"LoginIn\",\"result\":\"Error\"}\r\n"))
		return
	}

	// 搜索
	cmd := fmt.Sprintf("SELECT userName FROM %s WHERE userName=\"%s\" and password=\"%s\"",
		sql.SQLTableUser, jsonReslut["user"], jsonReslut["psw"])
	var user string
	// 从数据库中读取。
	err := sql.GetInstance().Get(cmd, func(param ...interface{}) {
	}, &user)

	if err != nil {
		client.Write([]byte("{\"type\":\"LoginIn\",\"result\":\"Error\"}\r\n"))
	} else {
		client.Write([]byte("{\"type\":\"LoginIn\",\"result\":\"OK\"}\r\n"))
	}
}

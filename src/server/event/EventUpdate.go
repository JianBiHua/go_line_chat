package event

import (
	"GoWorkspace/go_line_chat/src/server/sql"
	"net"
)

// EventUpdate is update event
type EventUpdate struct {
	client net.Conn
}

// Parse is parse update json string, and return result string
func (a *EventUpdate) Parse(jsonReslut map[string]interface{}, client net.Conn) {
	//
	a.client = client

	// 更新数据库
	var cmd = "UPDATE " + sql.SQLTableUser + " SET "
	// cmd := fmt.Sprintf("UPDATE UserTable1 SET userName=\"user2\" WHERE id=1",
	// 	sql.SQLTableUser, jsonReslut["user"], jsonReslut["psw"])
	// 更新密码
	if jsonReslut["psw"] != nil {
		cmd += " password=\"" + jsonReslut["psw"].(string) + "\""
	}
	// 更新密码
	// if jsonReslut["psw"] != nil {
	// 	cmd += " password=\"" + jsonReslut["psw"].(string) + "\""
	// }
	// if jsonReslut["psw"] != nil {
	// 	cmd += " password=\"" + jsonReslut["psw"].(string) + "\""
	// }

	cmd += " where userName = \"" + jsonReslut["user"].(string) + "\""
	sql.GetInstance().UpdateOrDelete(cmd)
	client.Write([]byte("{\"type\":\"Update\",\"result\":\"OK\"}\r\n"))
}

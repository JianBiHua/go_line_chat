package event

import (
	"GoWorkspace/go_line_chat/src/server/configer"
	"GoWorkspace/go_line_chat/src/server/sql"
	"fmt"
	"net"
	"time"
)

// EventLoginOut is login out event
type EventLoginOut struct {
	client net.Conn
}

// Parse is parse login out json string, and return result string
func (a *EventLoginOut) Parse(jsonReslut map[string]interface{}, client net.Conn) {
	//
	a.client = client

	// 获取用户名
	var userName = configer.ClientMap.GetKey(client)
	if len(userName) == 0 {
		client.Write([]byte("{\"type\":\"Chat\",\"result\":\"user is not login in\"}\r\n"))
		return
	}

	// 退出成功
	client.Write([]byte("{\"type\":\"LoginOut\",\"result\":\"OK\"}\r\n"))

	// 刷新LastDate时间
	cmd := fmt.Sprintf("UPDATE %s SET lastDate=\"%s\" WHERE userName=\"%s\"",
		sql.SQLTableUser, time.Now().Format("2006-01-02 15:04:05"), userName)
	sql.GetInstance().UpdateOrDelete(cmd)

	// 删除登出用户的net.Conn
	configer.ClientMap.Delete(client)
}

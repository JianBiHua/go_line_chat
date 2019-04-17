package event

import (
	"GoWorkspace/go_line_chat/src/common/model"
	"GoWorkspace/go_line_chat/src/server/configer"
	"GoWorkspace/go_line_chat/src/server/sql"
	"fmt"
	"net"
)

// EventFriendDel is sign in event
type EventFriendDel struct {
	client net.Conn
}

// 请求消息:
// {
//	"type" : "Del",
//	"user" : "user1"
// }

// Parse is parse del friend json string, and return result string
func (a *EventFriendDel) Parse(jsonReslut map[string]interface{}, client net.Conn) {
	//
	a.client = client

	// 如果不存在对端用户名，则打印错误
	if jsonReslut["user"] == nil {
		client.Write([]byte("{\"type\":\"Del\",\"result\":\"user not login in\"}\r\n"))
		return
	}

	// 获取用户名
	var userName = configer.ClientMap.GetKey(client)
	if len(userName) == 0 {
		client.Write([]byte("{\"type\":\"Del\",\"result\":\"Error\"}\r\n"))
		return
	}

	var user = new(model.User)
	// 先搜索是否存在
	cmd := fmt.Sprintf("SELECT userName FROM %s WHERE userName=\"%s\"",
		sql.SQLTableUser, jsonReslut["user"])

	err := sql.GetInstance().Get(cmd, func(param ...interface{}) {
	}, &user.UserName)

	// 没有该用户
	if err != nil {
		client.Write([]byte("{\"type\":\"Del\",\"result\":\"user2 is not exist\"}\r\n"))
		return
	}

	// 从好友列表删除
	cmd = fmt.Sprintf("DELETE FROM %s WHERE userName=\"%s\" and userName2=\"%s\"",
		sql.SQLTableFriends, userName, jsonReslut["user"].(string))
	_, err = sql.GetInstance().UpdateOrDelete(cmd)
	if err != nil {
		client.Write([]byte("{\"type\":\"Del\",\"result\":\"is not friends\"}\r\n"))
		return
	}

	// 告诉用户删除成功
	client.Write([]byte("{\"type\":\"Del\",\"result\":\"OK\"}\r\n"))

	// 告诉对端，我删除你了
	client2, ok := configer.ClientMap.Load(jsonReslut["user"].(string))
	if ok {
		client2.(net.Conn).Write([]byte("{\"type\":\"Del\",\"user\":\"" + userName + "\"}\r\n"))
	}
}

package event

import (
	"GoWorkspace/go_line_chat/src/server/configer"
	"GoWorkspace/go_line_chat/src/server/sql"
	"encoding/json"
	"fmt"
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

	// 获取用户名
	var userName = configer.ClientMap.GetKey(client)
	if len(userName) == 0 {
		client.Write([]byte("{\"type\":\"Chat\",\"result\":\"user is not login in\"}\r\n"))
		return
	}

	// 判断组是否存在,
	var ID int64
	cmd := fmt.Sprintf("SELECT id FROM %s WHERE id=%v",
		sql.SQLTableChatGroup, jsonReslut["group"])

	err := sql.GetInstance().Get(cmd, func(param ...interface{}) {
	}, &ID)

	if err != nil {
		// 该组已删除
		client.Write([]byte("{\"type\":\"Chat\",\"result\":\"group not found\"}\r\n"))
		return
	}

	// 开始找成员表.
	cmd = fmt.Sprintf("SELECT userName FROM %s WHERE groupId=%v",
		sql.SQLTableChatGroupMember, jsonReslut["group"])

	var tmpUserName string
	var users []string
	err = sql.GetInstance().Get(cmd, func(param ...interface{}) {
		// 只添加组员，不包括自己
		if userName != tmpUserName {
			users = append(users, tmpUserName)
		}
	}, &tmpUserName)

	// 发送消息给已登录的用户
	for _, name := range users {
		conn, ok := configer.ClientMap.Load(name)
		if !ok {
			continue
		}

		// 找到在线用户, 并发送消息给对端用户。
		result, _ := json.Marshal(jsonReslut)
		conn.(net.Conn).Write(result)
	}

	// 回发消息，发送成功
	client.Write([]byte("{\"type\":\"Chat\", \"result\": \"OK\"}"))

	// 将数据保存到ChatMsg数据库中.
	cmd = fmt.Sprintf("INSERT INTO %s(userName, groupId, comment, type) values(\"%v\",%v,\"%v\",%v)",
		sql.SQLTableChatMsg, userName, jsonReslut["group"], jsonReslut["msg"], jsonReslut["msgType"])
	sql.GetInstance().Insert(cmd)
}

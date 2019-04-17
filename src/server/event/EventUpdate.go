package event

import (
	"GoWorkspace/go_line_chat/src/server/configer"
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

	// 获取用户名
	var userName = configer.ClientMap.GetKey(client)
	if len(userName) == 0 {
		client.Write([]byte("{\"type\":\"EventChatGroupGet\",\"result\":\"user not login in\"}\r\n"))
		return
	}

	// 更新数据库
	var cmd = "UPDATE " + sql.SQLTableUser + " SET "
	var param = ""
	// cmd := fmt.Sprintf("UPDATE UserTable1 SET userName=\"user2\" WHERE id=1",
	// 	sql.SQLTableUser, jsonReslut["user"], jsonReslut["psw"])
	// 更新密码
	if jsonReslut["psw"] != nil {
		param += " password=\"" + jsonReslut["psw"].(string) + "\""
	}
	// 修改名字
	if jsonReslut["name"] != nil {
		if len(param) != 0 {
			param += ","
		}

		param += "name=\"" + jsonReslut["name"].(string) + "\""
	}
	//修改年龄
	if jsonReslut["age"] != nil {
		if len(param) != 0 {
			param += ","
		}

		param += "age=\"" + jsonReslut["age"].(string) + "\""
	}
	// 修改性别
	if jsonReslut["sex"] != nil {
		if len(param) != 0 {
			param += ","
		}

		param += "sex=" + jsonReslut["sex"].(string)
	}
	// 修改个性签名。
	if jsonReslut["signature"] != nil {
		if len(param) != 0 {
			param += ","
		}

		param += "signature=\"" + jsonReslut["signature"].(string) + "\""
	}

	// 最终的命令
	cmd += param
	cmd += " where userName = \"" + userName + "\""
	_, err := sql.GetInstance().UpdateOrDelete(cmd)

	if err == nil {
		client.Write([]byte("{\"type\":\"Update\",\"result\":\"OK\"}\r\n"))
	} else {
		client.Write([]byte("{\"type\":\"Update\",\"result\":\"update error\"}\r\n"))
	}
}

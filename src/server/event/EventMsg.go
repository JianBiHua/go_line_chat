package event

import (
	"GoWorkspace/go_line_chat/src/common/model"
	"GoWorkspace/go_line_chat/src/server/configer"
	"GoWorkspace/go_line_chat/src/server/sql"
	"GoWorkspace/go_line_chat/src/server/toolunit"
	"encoding/json"
	"fmt"
	"net"
)

// EventMsg is sign in event
type EventMsg struct {
	client net.Conn
}

// Parse is parse msg json string, and return result string
func (a *EventMsg) Parse(jsonReslut map[string]interface{}, client net.Conn) {
	//
	a.client = client

	// 获取用户名
	var userName = configer.ClientMap.GetKey(client)
	if len(userName) == 0 {
		client.Write([]byte("{\"type\":\"Chat\",\"result\":\"user is not login in\"}\r\n"))
		return
	}

	// 获取
	cmd := fmt.Sprintf("select lastDate from %s where userName=\"%s\"",
		sql.SQLTableUser, userName)

	var lastDate string
	sql.GetInstance().Get(cmd, func(param ...interface{}) {
	}, &lastDate)

	// 获取最后一次登录到这此登录的所有消息
	// cmd = fmt.Sprintf("select * from %s where groupId=%v and sendDate>\"%v\"",
	// 	sql.SQLTableChatMsg, jsonReslut["group"], lastDate)

	// 获取所有消息
	cmd = fmt.Sprintf("select * from %s where groupId=%v",
		sql.SQLTableChatMsg, jsonReslut["group"])

	var msg model.UserChatMessage
	var msgList []model.UserChatMessage
	sql.GetInstance().Get(cmd, func(param ...interface{}) {
		var tmp model.UserChatMessage
		toolunit.Copy2(&tmp, msg)
		msgList = append(msgList, tmp)
	}, &msg.ID, &msg.UserName, &msg.GroupID, &msg.SendDate, &msg.Comment, &msg.MsgType)

	//
	if len(msgList) > 0 {
		result, _ := json.Marshal(msgList)
		result = []byte(fmt.Sprintf("{\"type\" : \"Chat\", msgList : %s}", string(result)))
		client.Write(result)
	} else {
		client.Write([]byte("{\"type\":\"LoginOut\",\"result\":\"not msg\"}\r\n"))
	}
}

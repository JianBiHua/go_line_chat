package event

import (
	"GoWorkspace/go_line_chat/src/common/model"
	"GoWorkspace/go_line_chat/src/server/configer"
	"GoWorkspace/go_line_chat/src/server/sql"
	"fmt"
	"net"
)

// EventFriendAdd is sign in event
type EventFriendAdd struct {
	client net.Conn
}

// Parse is parse add friends json string, and return result string
func (a *EventFriendAdd) Parse(jsonReslut map[string]interface{}, client net.Conn) {
	//
	a.client = client

	// 将加好友信息保存到消息数组中，并且发送消息给对端。
	// MSGFriendAdd 添加好友
	//
	// 请求消息:
	// {
	//	"type" : "Add",
	//	"user" : "user2"  //对端用户名
	// }

	// 如果不存在对端用户名，则打印错误
	if jsonReslut["user"] == nil {
		client.Write([]byte("{\"type\":\"Add\",\"result\":\"Error\"}\r\n"))
		return
	}

	// 获取用户名
	var userName = configer.ClientMap.GetKey(client)
	if len(userName) == 0 {
		client.Write([]byte("{\"type\":\"Add\",\"result\":\"Error\"}\r\n"))
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
		client.Write([]byte("{\"type\":\"Add\",\"result\":\"User is not Exsit\"}\r\n"))
		return
	}

	// 判断是否是好友了
	cmd = fmt.Sprintf("SELECT userName FROM %s WHERE (userName=\"%s\" and userName2=\"%s\") or (userName=\"%s\" and userName2=\"%s\")",
		sql.SQLTableFriends, userName, jsonReslut["user"], jsonReslut["user"], userName)

	err = sql.GetInstance().Get(cmd, func(param ...interface{}) {
	}, &user.UserName)

	// 已经是好友了，不要重复添加。
	if err == nil {
		client.Write([]byte("{\"type\":\"Add\",\"result\":\"user is Friend!\"}\r\n"))
		return
	}

	// 偷偷懒，直接添加到好友列表中
	cmd = fmt.Sprintf("INSERT INTO %s(userName, userName2) values(\"%s\",\"%v\")",
		sql.SQLTableFriends, userName, jsonReslut["user"])
	//
	_, err = sql.GetInstance().Insert(cmd)
	if err != nil {
		// 告诉用户添加失败
		client.Write([]byte("{\"type\":\"Add\",\"result\":\"insert failed!\"}\r\n"))
		return
	}

	// 告诉用户添加成功
	client.Write([]byte("{\"type\":\"Add\",\"result\":\"OK\"}\r\n"))

	// 告诉对端，我添加你为好友了
	client2, ok := configer.ClientMap.Load(jsonReslut["user"].(string))
	if ok {
		client2.(net.Conn).Write([]byte("{\"type\":\"Add\",\"user\":\"" + userName + "\"}\r\n"))
	}

	// 插入到消息数据库
	// cmd = fmt.Sprintf("INSERT INTO %s(userName, userName2, type) values(\"%s\",\"%v\",%d)",
	// 	sql.SQLTableChatMsg, userName, jsonReslut["user"], message.MSGTypeAddFriend)
	// _, err = sql.GetInstance().Insert(cmd)
	// if err != nil {
	// 	client.Write([]byte("{\"type\":\"Add\",\"result\":\"Error\"}\r\n"))
	// } else {
	// 	client.Write([]byte("{\"type\":\"Add\",\"result\":\"OK\"}\r\n"))
	// }
}

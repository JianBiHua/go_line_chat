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

// EventFriendGet is sign in event
type EventFriendGet struct {
	client net.Conn
}

// Parse is parse get friends json string, and return result string
func (a *EventFriendGet) Parse(jsonReslut map[string]interface{}, client net.Conn) {
	//
	a.client = client

	var userName = configer.ClientMap.GetKey(client)
	if len(userName) == 0 {
		client.Write([]byte("{\"type\":\"EventChatGroupGet\",\"result\":\"user not login in\"}\r\n"))
		return
	}

	//
	var users []*model.User

	// 搜索好友
	cmd := fmt.Sprintf("SELECT userName,userName2 FROM %s WHERE userName=\"%s\" or userName2=\"%s\"",
		sql.SQLTableFriends, userName, userName)

	var userName1, userName2 string
	err := sql.GetInstance().Get(cmd, func(param ...interface{}) {
		user := new(model.User)

		if userName1 != userName {
			user.UserName = userName1
		} else {
			user.UserName = userName2
		}

		users = append(users, user)
	}, &userName1, &userName2)

	if err != nil {
		client.Write([]byte("{\"type\":\"FriendGet\",\"result\":\"user no have friends\"}\r\n"))
		return
	}

	// 获取好友详细信息。
	user := new(model.User)

	// 根据用户名读取组员详细信息
	for _, value := range users {
		cmd = fmt.Sprintf("SELECT * FROM %s WHERE userName=\"%s\"",
			sql.SQLTableUser, value.UserName)

		err = sql.GetInstance().Get(cmd, func(param ...interface{}) {
			user2 := new(model.User)
			// 深拷贝
			toolunit.Copy(user2, user)
			// 追加到切片中。
			users = append(users, user2)
		}, &user.Id, &user.UserName, &user.Password, &user.Sex, &user.Name,
			&user.Age, &user.Icon, &user.Signature, &user.CreateDate, &user.LastDate, &user.Status)

		if err != nil {
			client.Write([]byte("{\"type\":\"FriendGet\",\"result\":\"" + err.Error() + "\"}\r\n"))
			return
		}
	}

	// 发送好友列表给目标
	result, _ := json.Marshal(users)
	client.Write(result)
}

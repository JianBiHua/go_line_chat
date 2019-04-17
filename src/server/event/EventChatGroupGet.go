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

// EventChatGroupGet is new Chat Group event
type EventChatGroupGet struct {
	client net.Conn
}

// Parse is parse get Chat group json string, and return result string
// 这里联动查询了三个数据库，才得到所有的结果。
func (a *EventChatGroupGet) Parse(jsonReslut map[string]interface{}, client net.Conn) {
	//
	a.client = client

	// 获取用户名
	var userName = configer.ClientMap.GetKey(client)
	if len(userName) == 0 {
		client.Write([]byte("{\"type\":\"EventChatGroupGet\",\"result\":\"user is not login in\"}\r\n"))
		return
	}

	//
	var groups []*model.UserChatGroup
	// 临时组变量
	group := new(model.UserChatGroup)

	// 搜索
	cmd := fmt.Sprintf("SELECT * FROM %s WHERE userName=\"%s\"",
		sql.SQLTableChatGroup, jsonReslut["user"])

	// 从数据库中读取组信息。
	err := sql.GetInstance().Get(cmd, func(param ...interface{}) {
		group2 := new(model.UserChatGroup)
		// 深拷贝
		toolunit.Copy(group2, group)
		// 追加到切片中。
		groups = append(groups, group2)
	}, &group.Id, &group.Name, &group.Signature, &group.Max, &group.GroupType, &group.Icon, &group.CreateDate)

	// 根据数组读取组用户列表。
	for _, group := range groups {
		cmd := fmt.Sprintf("SELECT userName,joinDate FROM %s WHERE id=%d",
			sql.SQLTableChatGroupMember, group.Id)
		user := new(model.User)

		err := sql.GetInstance().Get(cmd, func(param ...interface{}) {
			user2 := new(model.User)
			// 深拷贝
			toolunit.Copy(user2, user)
			// 追加到切片中。
			group.Member = append(group.Member, user2)
		}, &user.UserName, &user.JoinDate)

		// 根据用户名读取组员详细信息
		if err == nil {
			cmd = fmt.Sprintf("SELECT * FROM %s WHERE userName=\"%s\"",
				sql.SQLTableUser, user.UserName)

			sql.GetInstance().Get(cmd, func(param ...interface{}) {
				user2 := new(model.User)
				// 深拷贝
				toolunit.Copy(user2, user)
				// 追加到切片中。
				group.Member = append(group.Member, user2)
			}, &user.Id, &user.UserName, &user.Password, &user.Sex, &user.Name,
				&user.Age, &user.Icon, &user.Signature, &user.CreateDate, &user.LastDate, &user.Status)
		}
	}

	// 将获取到的信息转换成json，发给客户端.
	result, err := json.Marshal(groups)
	if err != nil {
		client.Write(result)
	} else {
		client.Write([]byte("{\"type\":\"EventChatGroupGet\",\"result\":\"group not found\"}\r\n"))
	}
}

package event

import (
	"GoWorkspace/go_line_chat/src/common/message"
	"GoWorkspace/go_line_chat/src/common/model"
	"GoWorkspace/go_line_chat/src/server/configer"
	"GoWorkspace/go_line_chat/src/server/sql"
	"encoding/json"
	"fmt"
	"net"
)

// EventChatGroupCreate is new Chat Group event
type EventChatGroupCreate struct {
	client net.Conn
}

// {
// 	"type" : "GroupCreate",
// 	"users" : ["user1", "user2"] //对端用户名组。
//  "groupType" : 0	// 组类型
// }
// Parse is parse create Chat group json string, and return result string
func (a *EventChatGroupCreate) Parse(jsonReslut map[string]interface{}, client net.Conn) {
	//
	a.client = client

	// 判断登陆
	var userName = configer.ClientMap.GetKey(client)
	if len(userName) == 0 {
		client.Write([]byte("{\"type\":\"GroupCreate\",\"result\":\"user is not login in\"}\r\n"))
		return
	}

	group := new(model.UserChatGroup)

	// 如果是一对一，则判断是否已创建过组了
	if jsonReslut["groupType"].(float64) == message.MSGGroupTypeOne2One ||
		jsonReslut["groupType"].(float64) == message.MSGGroupTypeOne2One2 {

		// 联合查询找出聊天组表中的项
		// 查找组员数据库中，groundId相同，并且只有两个， userName包含username并且username2的项
		// 1. 查找组员数据库中，groundId相同，并且只有两个
		// A. 这种获取有个问题，获取到的是所有与userName,userName2的聊天组
		// select * from ChatGroupMember where groupId in (select groupId from ChatGroupMember group by groupId having count(*)=2)
		// B. 获取到的是userName与userName2的聊天组
		// select * from ChatGroupMember where ChatGroupMember.groupId in (select groupId from ChatGroupMember group by ChatGroupMember.groupId having count(*)=2 and userName in ("test1", "test2"))
		// 2. userName包含username或者username2的项
		// select * from ChatGroupMember where ChatGroupMember.groupId in (select groupId from ChatGroupMember group by ChatGroupMember.groupId having count(*)=2) and (userName="test3" or userName="test4")
		// 3. 联合表ChatGroup查询出组信息,
		// select g.* from ChatGroup g,(select * from ChatGroupMember where ChatGroupMember.groupId in (select groupId from ChatGroupMember group by ChatGroupMember.groupId having count(*)=2 and userName in ("test1", "test2"))) where g.id=groupId and (type=0 or type=1)
		// 4. 取消相同数据。
		// select distinct * from (select g.* from ChatGroup g,(select * from ChatGroupMember m where m.groupId in (select groupId from ChatGroupMember gm group by gm.groupId having count(*)=2 and userName in ("test1", "test2"))) where g.id=groupId and (type=0 or type=1))

		cmd := fmt.Sprintf("select distinct * from (select g.* from %s g,(select * from %s m where m.groupId in (select groupId from %s gm group by gm.groupId having count(*)=2 and userName in (\"%v\", \"%v\"))) where g.id=groupId and (type=0 or type=1))",
			sql.SQLTableChatGroup, sql.SQLTableChatGroupMember, sql.SQLTableChatGroupMember, userName, jsonReslut["users"].([]interface{})[0])

		// 从数据库中读取组信息。
		// 只能查出一项，如果查出多个了，肯定就错了。
		sql.GetInstance().Get(cmd, func(param ...interface{}) {
		}, &group.ID, &group.Name, &group.Signature, &group.Max, &group.GroupType, &group.Icon, &group.CreateDate)
	}

	if group.ID == 0 {
		// 创建组，并且将组员信息加入。
		cmd := fmt.Sprintf("INSERT INTO %s(name, type) values(\"group\",%v)", sql.SQLTableChatGroup, jsonReslut["groupType"])
		groupID, _ := sql.GetInstance().Insert(cmd)

		// 获取组信息
		cmd = fmt.Sprintf("select * from %s where id=%d", sql.SQLTableChatGroup, groupID)
		sql.GetInstance().Get(cmd, func(param ...interface{}) {
		}, &group.ID, &group.Name, &group.Signature, &group.Max, &group.GroupType, &group.Icon, &group.CreateDate)

		// 插入组成员信息。
		for _, value := range jsonReslut["users"].([]interface{}) {
			cmd := fmt.Sprintf("INSERT INTO %s(groupId, userName) values(%d,\"%v\")", sql.SQLTableChatGroupMember, groupID, value)
			sql.GetInstance().Insert(cmd)
		}

		// 将自己插入到成员表
		cmd = fmt.Sprintf("INSERT INTO %s(groupId, userName) values(%d,\"%v\")", sql.SQLTableChatGroupMember, groupID, userName)
		sql.GetInstance().Insert(cmd)
	}

	// 返回组信息
	// 将获取到的信息转换成json，发给客户端.
	// 学习使用反射。
	result, err := json.Marshal(group)
	if err != nil {
		client.Write([]byte("{\"type\":\"GroupCreate\",\"result\":\"create group error!\"}\r\n"))
		return
	}

	result = []byte(fmt.Sprintf("{\"type\" : \"GroupCreate\", %s}", string(result)))
	client.Write(result)
}

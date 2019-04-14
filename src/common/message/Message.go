package message

// 消息反馈情况类型。
// 使用Json类型。
const (

	// MSGResultNone 消息结果，无信息
	MSGResultNone = "None"

	// MSGResultSuccess 消息结果，成功
	MSGResultSuccess = "OK"

	// MSGResultFailed 消息结果，失败
	MSGResultFailed = "Error"
)

// 聊天消息类型
// 使用Json类型。
const (

	// MSGSignIn 注册消息
	//
	// 请求消息:
	// {
	//	"type" : "SignIn",
	//	"user" : "user1",
	//  "psw" : "psw1"
	// }
	//
	// 返回消息
	// 成功时
	// {
	//	 "type" : "SignIn",
	// 	"result" : "Ok"
	// }
	// 失败时
	// {
	//	 "type" : "SignIn",
	// 	"result" : "Error"
	// }
	MSGSignIn = "SignIn"

	// MSGLoginIn 登陆
	//
	// 请求消息:
	// {
	//	"type" : "LoginIn",
	//	"user" : "user1",
	//  "psw" : "psw1"
	// }
	MSGLoginIn = "LoginIn"

	// MSGLoginOut 退出
	//
	// 请求消息:
	// {
	//	"type" : "LoginOut"
	// }
	MSGLoginOut = "LoginOut"

	// MSGFriendAdd 添加好友
	//
	// 请求消息:
	// {
	//	"type" : "Add",
	//	"user" : "user2"  //对端用户名
	// }
	MSGFriendAdd = "Add"

	// MSGFriendDel 删除好友
	//
	// 请求消息:
	// {
	//	"type" : "Del",
	//	"user" : "user2"  //对端用户名
	// }
	MSGFriendDel = "Del"

	// MSGNewChatGroup 建聊天组
	//
	// 请求消息:
	// {
	//	"type" : "Group",
	//	"users" : ["user1", "user2"] //对端用户名组。
	//  "isGroup" : 0	//是否是群组
	// }
	MSGNewChatGroup = "Group"

	// MSGChat 聊天，发送聊天信息。
	//
	// 请求消息:
	// {
	//	"type" : "Chat",
	//	"grop" : 1 组ID。
	//  "msg" : "我发来一条消息"   //发送的消息
	// }

	// 返回消息:
	// {
	// "type" : "Chat",
	// "result" : "Ok"
	// }
	MSGChat = "Chat"

	// MSGUpdate 更新用户信息。
	//
	// 请求消息:
	// {
	//	"type" : "Update",
	//	"name" : "张三"
	//  "age" : "1980/01/19" 	//年龄
	//  ...
	// }
	MSGUpdate = "Update"
)

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

	// MSGChatGroupCreate 建聊天组
	//
	// 请求消息:
	// {
	//	"type" : "GroupCreate",
	//	"users" : ["user1", "user2"] //对端用户名组。
	//  "groupType" : 0	//组类型
	// }
	MSGChatGroupCreate = "GroupCreate"

	// MSGChatGroupCreate 建聊天组
	//
	// 请求消息:
	// {
	//	"type" : "GroupGet"
	// }
	MSGChatGroupGet = "GroupGet"

	// MSGChat 聊天，发送聊天信息。
	//
	// 请求消息:
	// {
	//	"type" : "Chat",
	//	"group" : 1 组ID。
	//  "msg" : "我发来一条消息"   //发送的消息
	//  "msgType" : 0
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

	// MSGFriendGet 获取好友信息
	//
	// {
	//	"type" : "FriendGet"
	// }
	MSGFriendGet = "FriendGet"

	// MSGMsg 获取说有历史消息
	//
	// {
	//	"type" : "Msg"
	//	"group" : 0
	// }
	MSGMsg = "Msg"
)

// MSGType is msg type
type MSGType int

// msg type
const (
	// MSGTypeText 文字
	MSGTypeText = iota
	//
	MSGTypePicture
	//
	MSGTypeAudio
	//
	MSGTypeVideo
	//
	MSGTypeWebPage
	//
	MSGTypeAddFriend
	//
	MSGTypeAddFriendSuccess
)

// MSGGroupType is group type
type MSGGroupType int

// 一对一， 群(qq)， 讨论组(qq)，公众号, 临时聊天 ...等等
// 现只管MSGGroupTypeOne2One，MSGGroupTypeOne2One2，MSGGroupTypeGroup
const (
	// MSGGroupTypeOne2One 一对一
	MSGGroupTypeOne2One = iota

	// MSGGroupTypeOne2One2 一对一, 私聊，不是好友的情况
	MSGGroupTypeOne2One2

	// MSGGroupTypeGroup 群聊
	MSGGroupTypeGroup
)

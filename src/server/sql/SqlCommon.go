package sql

// sql names define
const (

	// 数据库名称
	SQLDataBaseName = "LineChat.db"

	// 用户表
	//
	// 用来保存用户的信息，用户名密码等.
	SQLTableUser = "ChatUser"

	// 聊天组表
	//
	// 用来保存用户聊天组的信息。
	SQLTableChatGroup = "ChatGroup"

	// 聊天组表
	//
	// 用来保存用户聊天组组员的信息。
	SQLTableChatGroupMember = "ChatGroupMember"

	// 聊天内容信息表
	//
	// 用来保存各组的聊天信息的表.
	SQLTableChatMsg = "ChatMsg"

	// 好友信息表
	//
	// 用来保存好友信息的表.
	SQLTableFriends = "ChatFriends"
)

// create database and tables sql command
const (
	// create User Table's sql command
	//
	// table member item info:
	// id :
	// user : 用户名
	// password : 密码
	// sex : 性别
	// name : 显示的名称
	// age : 年龄
	// icon : 头像路径
	// signature : 个性签名
	// joinDate : 创建日期时间
	// lastDate : 最后一次登陆日期时间
	// status : 当前状态, 0： 离线 1: 在线  2: 隐身 3: 免打扰 等等..
	SQLCommandUser = "CREATE TABLE IF NOT EXISTS " + SQLTableUser +
		`(
			"id" INTEGER PRIMARY KEY AUTOINCREMENT,
			"userName" varchar(30) unique,
			"password" varchar(30),
			"sex" int(2) default 0,
			"name" varchar(20) default "",
			"age" TIMESTAMP default (datetime('now', 'localtime')) ,
			"icon" varchar (200) default "",
			"signature" varchar (500) default "",
			"joinDate" TIMESTAMP default (datetime('now', 'localtime')),
			"lastDate" TIMESTAMP default (datetime('now', 'localtime')),
			"status" int(5) default 0
		); 
		` +
		"CREATE INDEX IF NOT EXISTS User ON " + SQLTableUser + "(userName);"

	// create User Chat Group Table's sql command
	//
	// table member item info:
	// id :
	// name : 组名称。
	// signature: 组介绍，个性签名
	// max: 最大个数
	// icon: 图标路径
	// createDate: 创建时间
	// type: 组类型，
	SQLCommandUserChatGroup = "CREATE TABLE IF NOT EXISTS " + SQLTableChatGroup +
		`(
			"id" INTEGER PRIMARY KEY AUTOINCREMENT,
			"name" varchar(20) default "",
			"signature" varchar(500) default "",
			"max" int(8) default 100,
			"type" int(8) default 0,
			"icon" varchar(200) default "",
			"createDate" TIMESTAMP default (datetime('now', 'localtime'))
		);
	` +
		"CREATE INDEX IF NOT EXISTS GroupType ON " + SQLTableChatGroup + "(type);"

	// create User Chat Group member Table's sql command
	//
	// table member item info:
	// id :
	// user : 用户名。
	// joinDate: 进入组时间。
	SQLCommandUserChatGroupMember = "CREATE TABLE IF NOT EXISTS " + SQLTableChatGroupMember +
		`(
			"id" INTEGER PRIMARY KEY AUTOINCREMENT,
			"groupId"  INTEGER,
			"userName" varchar(30) not null,
			"joinDate" TIMESTAMP default (datetime('now', 'localtime'))
		);
	` +
		"CREATE INDEX IF NOT EXISTS MemberIndex ON " + SQLTableChatGroupMember + "(userName, groupId);"

	// create User Chat old msg Table's sql command
	//
	// table member item info:
	// id :
	// user : 发送消息的用户。
	// sendDate: 发送的时间。
	// comment: 发送的内容, 可能是文件，或者图片路径或者网页路径等。
	// type: 消息类型: 0: text 1: picture 2: audio 3: video 4: web page，5: 添加好友
	SQLCommandUserChatMsg = "CREATE TABLE IF NOT EXISTS " + SQLTableChatMsg +
		`(
			"id" INTEGER PRIMARY KEY AUTOINCREMENT,
			"userName" varchar(30) not null,
			"groupId" int default 0,
			"sendDate" TIMESTAMP default (datetime('now', 'localtime')),
			"comment" TEXT default "",
			"type" int default 0
		);` +
		"CREATE INDEX IF NOT EXISTS MsgIndex ON " + SQLTableChatMsg + "(userName, groupId);"

	// create User Friends Table's sql command
	//
	// table friends info:
	// id :
	// user1: user 1
	// user2: user 2
	// joniDate: 成为好友的时间。
	SQLCommandFriends = "CREATE TABLE IF NOT EXISTS " + SQLTableFriends +
		`(
			"id" INTEGER PRIMARY KEY AUTOINCREMENT,
			"userName" varchar(30) not null,
			"userName2" varchar(30) not null,
			"joniDate" TIMESTAMP default (datetime('now', 'localtime'))
		);
	` +
		// 创建联合唯一索引
		"CREATE unique INDEX IF NOT EXISTS FriendsUser ON " + SQLTableFriends + "(userName, userName2);"
)

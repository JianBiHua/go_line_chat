go_line_chat
==========
1. 并发聊天室服务器端程序，客户端程序[go程序/android程序]
2. 正在研发中...

### Environment
开发语言: `GO 1.12.1` \
开发工具: `VSCode` \
数据库: `sqlite3`

### Course
教程地址: https://blog.csdn.net/dkaily1314/article/details/88825075

### Installation
```bash

$ cd go_line_chat
$ ./Build.sh

```
>
> 将生成可执行文件到go_line_chat/Bin下, 包括mac/windows/linux程序,现只有windows可以看到图标
>

### Clow Chart
![Image text](https://github.com/JianBiHua/go_line_chat/blob/master/resources/%E6%B5%81%E7%A8%8B%E5%9B%BE_v1_0.png)

### Event List
``` json
// 注册
{
"type":"SignIn",
"user":"test1",
"psw":"pasw"
}

// 登陆
{
"type":"LoginIn",
"user":"test1",
"psw":"pasw"
}

// 登出
{
"type":"LoginOut"
}

// 获取好友列表
{
"type" : "FriendGet"
}

// 添加好友
{
"type" : "Add",
"user" : "test2"
}

// 删除好友
{
"type" : "Del",
"user" : "test1"
}

// 获取聊天组
{
"type" : "GroupGet"
}

// 创建组
{
	"type" : "GroupCreate",
	"users" :["test2"],
	"groupType" : 0
}

// 发送消息
{
	"type" : "Chat",
	"group" : 2,
 	"msg" : "我发来一条消息",
 	"msgType" : 0
}

// 获取组2的消息
{
	"type" : "Msg",
	"group" : 2
}
```

### Libraries
| name | go get | Description |
|:---:|:---|:---|
| go-sqlite3 | go get github.com/mattn/go-sqlite3| sqlite3数据库 |
| rsrc | go get github.com/akavel/rsrc| 为windows生成图标文件 |
| upx | https://github.com/upx/upx/releases | 壳工具 |


### Author

简笔画

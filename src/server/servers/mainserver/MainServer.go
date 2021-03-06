// Package mainserver is main server
package mainserver

import (
	"GoWorkspace/go_line_chat/src/server/chatlog"
	"GoWorkspace/go_line_chat/src/server/configer"
	"GoWorkspace/go_line_chat/src/server/event"
	"GoWorkspace/go_line_chat/src/server/servers"
	"GoWorkspace/go_line_chat/src/server/sql"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

// MainServer 主服务器结构体.
type MainServer struct {
	// 服务器端口号
	port int
	// 主服务监听指针。
	listener net.Listener
	// 事件工具
	eventFactory *event.EventFactory
}

// StartMainServer 静态函数，启动主服务器服务。
func StartMainServer(p int) *MainServer {
	var server = new(MainServer)
	// 设置端口。
	server.port = p
	//
	server.eventFactory = event.NewEventFactory()
	// 启动服务。
	server.Start()

	return server
}

func (mainServer *MainServer) log(msg string) {
	chatlog.Append(chatlog.LOGMAINSERVER, msg)
}

// Start 开始服务
func (mainServer *MainServer) Start() servers.ServerErrorType {
	mainServer.log(fmt.Sprintf("启动主服务器[端口: %d]\r\n", mainServer.port))
	var ip = ""
	ip = fmt.Sprintf("127.0.0.1:%d", mainServer.port)
	//1. 监听
	var err error
	mainServer.listener, err = net.Listen("tcp", ip)
	if err != nil {
		mainServer.log(fmt.Sprintf("net.Listen: %s", err))
		return servers.ServerErrorListen
	}

	// 启动监听
	go mainServer.startListen()

	return servers.ServerErrorNone
}

// Stop 停止服务
func (mainServer *MainServer) Stop() {
	mainServer.log("停止主服务器\r\n")
}

// Restart 重启服务
func (mainServer *MainServer) Restart() {
	mainServer.Stop()
	mainServer.Start()
}

// startListen 开启监听
func (mainServer *MainServer) startListen() {
	mainServer.log("主服务器开始监听\r\n")
	for {
		// 客户端连接。
		client, err := mainServer.listener.Accept()
		if err != nil {
			mainServer.log("conn fail ...")
		} else {
			mainServer.log(fmt.Sprintf("客户端连接: %s", client.RemoteAddr().String()))

			// 当客户连接时，启动一个线程，处理线程的消息。
			go mainServer.onClientConnect(client)
		}
	}
}

// onClientConnect 当有客户端链接时，处理。
func (mainServer *MainServer) onClientConnect(client net.Conn) {
	for {
		data := make([]byte, 1024)
		// 读取数据
		msgLength, err := client.Read(data)

		if msgLength == 0 || err != nil {
			mainServer.log(fmt.Sprintf("客户端断开: %s", client.RemoteAddr().String()))
			//
			mainServer.onClientDisConnect(client)
			return
		}

		//
		msg := string(data[0:msgLength])
		go mainServer.onReadClientData(client, msg)
	}
}

// 当客户端断开时。
func (mainServer *MainServer) onClientDisConnect(client net.Conn) {
	// 从ClientMap移除
	configer.ClientMap.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		if v == client {
			// 如果是登陆用户断开连接时
			userName := configer.ClientMap.GetKey(v)
			if len(userName) > 0 {
				// 刷新LastDate时间
				cmd := fmt.Sprintf("UPDATE %s SET lastDate=\"%s\" WHERE userName=\"%s\"",
					sql.SQLTableUser, time.Now().Format("2006-01-02 15:04:05"), userName)
				sql.GetInstance().UpdateOrDelete(cmd)
			}

			configer.ClientMap.Delete(k)
			return false
		}
		return true
	})
}

// 当读到客户端数据时。
func (mainServer *MainServer) onReadClientData(client net.Conn, msg string) {
	// json数据。
	//var jsonReslut map[string]interface{}
	jsonReslut := make(map[string]interface{})

	// m := make(map[string]interface{})
	// m["type"] = "GroupType"
	// m["users"] = []string{"123", "234"}
	// m["groupType"] = 0

	// res, _ := json.Marshal(m)
	// fmt.Println(string(res))

	// err1 := json.Unmarshal([]byte(res), &jsonReslut)
	// fmt.Printf("err1: %v, %v", err1, jsonReslut)

	// 解析Json数据
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	// jsonReslut, ok := gjson.Parse(msg).Value().(map[string]interface{})
	// if !ok {
	// 	// not a map
	// 	return
	// }

	// fmt.Printf("jsonReslut[users]= %v", jsonReslut["users"])
	err := json.Unmarshal([]byte(msg), &jsonReslut)
	if err != nil {
		return
	}

	// 判断json数据是否合法
	var eventType = jsonReslut["type"]
	if eventType == nil {
		return
	}

	// 如果是合法的类型，则进入分发消息阶段。
	mainServer.eventFactory.Parse(jsonReslut, client)
}

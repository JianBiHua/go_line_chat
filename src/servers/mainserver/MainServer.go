// Package mainserver is main server
package mainserver

import (
	"GoWorkspace/go_line_chat/src/chatlog"
	"GoWorkspace/go_line_chat/src/servers"
	"fmt"
	"net"
)

// MainServer 主服务器结构体.
type MainServer struct {
	// 服务器端口号
	port int
	// 主服务监听指针。
	listener net.Listener
}

// StartMainServer 静态函数，启动主服务器服务。
func StartMainServer (p int) *MainServer {
	var server = new (MainServer)
	// 设置端口。
	server.port = p
	// 启动服务。
	server.Start ()

	return server
}

func (mainServer *MainServer) log (msg string)  {
	chatlog.Append(chatlog.LOGMAINSERVER, msg)
}

// Start 开始服务
func (mainServer *MainServer) Start () servers.ServerErrorType {
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
	go mainServer.startListen ()

	return servers.ServerErrorNone
}

// Stop 停止服务
func (mainServer *MainServer) Stop ()  {
	mainServer.log("停止主服务器\r\n")
}

// Restart 重启服务
func (mainServer *MainServer) Restart ()  {
	mainServer.Stop()
	mainServer.Start()
}

// startListen 开启监听
func (mainServer *MainServer) startListen ()  {
	mainServer.log("主服务器开始监听\r\n")
	for {
		// 客户端连接。
		client, err := mainServer.listener.Accept()
		if err != nil {
			mainServer.log("conn fail ...")
		} else {
			mainServer.log( fmt.Sprintf("客户端连接: %s", client.RemoteAddr().String()))

			// 当客户连接时，启动一个线程，处理线程的消息。
			go mainServer.onClientConnect (client)
		}
	}
}

// onClientConnect 当有客户端链接时，处理。
func (mainServer *MainServer) onClientConnect (client net.Conn)  {
	for {
		data := make([]byte, 255)
		// 读取数据
		msgLength, err := client.Read(data)

		if msgLength == 0 || err != nil {
			mainServer.log(fmt.Sprintf("客户端断开: %s", client.RemoteAddr().String()))
			return
		}

		//
		msg := string(data[0:msgLength])
		mainServer.log(fmt.Sprintf("Msg: %s", msg))
	}
}


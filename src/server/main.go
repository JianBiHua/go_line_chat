package main

import (
	"GoWorkspace/go_line_chat/src/server/chatlog"
	"GoWorkspace/go_line_chat/src/server/configer"
	"GoWorkspace/go_line_chat/src/server/servers/mainserver"
	"bufio"
	"flag"
	"os"
)

// 打印输出日志，
func log(msg string) {
	chatlog.Append(chatlog.LOGMAIN, msg)
}

// 显示菜单，
func showMenu() {
	log("===========命令菜单======")
	// 打印菜单像
	log("users--显示所有的注册用户")
	log("onlines--显示所有的在线用户")
	log("states--显示所有服务器的状态")
	log("ports--显示所有服务器的端口")
	log("restart all--重启所有服务器")
	log("restart login--重启登陆服务器")
	log("restart file--重启文件服务器")
	log("restart chat--重启聊天服务器")
	log("logmode X--修改日志模式[0-5]")
}

// 使用方法:
//
// 生成应用:
// go build -o LineChat
//
// 执行应用:
// LineChat -port 8889
func main() {
	// 通过命令行, 获取服务的端口号
	var port int
	flag.IntVar(&port, "port", 8888, "set main server port")
	flag.Parse()

	// 启动Server
	mainserver.StartMainServer(port)

	// 设置日志模式。
	configer.CurrentLogMode = chatlog.LOGMAIN

	// 显示菜单
	showMenu()

	// 循环读取输入
	input := bufio.NewScanner(os.Stdin) //初始化一个扫表对象
	log("输入命令: ")
	for input.Scan() {
		cmd := input.Text() //把输入内容转换为字符串
		switch cmd {
		case "menu":
			showMenu()
		case "users":
			log("--显示所有的注册用户--")
		case "onlines":
			log("--显示所有的在线用户--")
		case "states":
			log("--显示所有服务器的状态--")
		case "ports":
			log("--显示所有服务器的端口--")
		case "restart all":
			log("--显示所有服务器的端口--")
		case "restart login":
			log("--重启登陆服务器--")
		case "restart file":
			log("--重启文件服务器--")
		case "restart chat":
			log("--重启聊天服务器--")
		case "logmode 0":
			log("--修改日志模式为: 主程序模式--")
			configer.CurrentLogMode = chatlog.LOGMAIN
		case "logmode 1":
			log("--修改日志模式为: 主服务器模式--")
			configer.CurrentLogMode = chatlog.LOGMAINSERVER
		default:
			log("命令错误: " + cmd)
		}

		log("===输入命令: ")
	}
}

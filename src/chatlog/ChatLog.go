// Package chatlog chat log
package chatlog

import (
	"GoWorkspace/go_line_chat/src/configer"
	"GoWorkspace/go_line_chat/src/toolunit"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func getPath (logType LogType) string {
	// 获取日志文件路径。
	path, _ := toolunit.GetPathInstance().GetLogPath(GetLogName(logType))
	return path
}

// Append 将消息保存到指定的日志文件下，如果没有则创建该日志文件。
// 参数
// logType: 日志类型
// 返回
// 无
// 转换文件格式的函数
//decoder := mahonia.NewDecoder("gbk")
//f.gbkFile = decoder.NewReader(f.file)
func Append (logType LogType, msg string)  {
	// 获取应用目录
	path := getPath(logType)
	// 打开一个追加文件
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		//
		return
	}

	// 延迟关闭
	defer file.Close()

	// 消息加当前时间，
	var content = time.Now().Format("2006-01-02 15:04:05") +" "+ msg+"\r\n"

	// 打印日志。
	if configer.CurrentLogMode == int(logType) {
		fmt.Print(content)
	}

	// 写日志
	// 每行为一条信息。
	// 格式化输出时间。
	file.WriteString(content)
}

// Clear 清除指定的消息，删除该日志文件。
// 参数
// logType: 日志类型
// 返回
// 无
func Clear (logType LogType) {
	// 删除指定类型的目录文件。
	path := getPath(logType)
	os.Remove(path)
}

// ClearAll 直接删除日志目录下的所有文件
func ClearAll ()  {
	// 直接删除日志目录。
	log,_ := toolunit.GetPathInstance().GetLogPath("")
	os.RemoveAll(log)
}

// Read 读取日志消息，从日志文件
// 参数
// logType: 日志类型
// 返回
// string: 日志文件中的所有内容
// error: 读文件出的错误
func Read (logType LogType) (string, error)  {
	// 获取应用目录
	path := getPath(logType)
	// 打开一个追加文件
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	// 读取数据。
	bytes, err := ioutil.ReadAll(f)

	return string(bytes), err
}

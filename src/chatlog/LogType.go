// Package chatlog is chat log
package chatlog

import (
	"fmt"
)

// LogType log type
type LogType int

const (
	// LOGMAIN is main log
	LOGMAIN = iota
	// LOGMAINSERVER is main server log
	LOGMAINSERVER
	// LOGLOGIN is login server log
	LOGLOGIN
	// LOGFILE is file server log
	LOGFILE
	// LOGCHAT is chat server log
	LOGCHAT

	// LOGUSER 用户自定义消息类型。
	LOGUSER = 100
)

// GetLogName 获取LogType对应的文件名
// 参数:
// t: 日志类型。
// 返回:
//	string 日志文件名
func GetLogName (t LogType) string  {
	switch t {
	case LOGMAIN:
		return "main.log"
	case LOGMAINSERVER:
		return "mainserver.log"
	}

	return fmt.Sprintf("%d.log", t)
}
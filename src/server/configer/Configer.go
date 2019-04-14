package configer

import "sync"

// 当前的日志模式
var CurrentLogMode int

// 所有登陆用户
// [string]tcp.Conn
// [用户名]clientConn
var ClientMap sync.Map

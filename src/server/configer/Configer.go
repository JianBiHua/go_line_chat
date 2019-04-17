package configer

import (
	"sync"
)

// CurrentLogMode 当前的日志模式
var CurrentLogMode int

// SyncMap 将sync.Map定义为SyncMap类型
// 为其扩展功能
//
// 这种方法不行...，只能匿名
// type SyncMap sync.Map
type SyncMap struct {
	sync.Map
}

// ClientMap 所有登陆用户
// [string] net.Conn
// [用户名]  clientConn
var ClientMap SyncMap

// GetKey is get username form ClientMap
//
// input:
//	conn: one value
func (m SyncMap) GetKey(value interface{}) string {
	var key = ""

	// 遍历所有sync.Map中的键值对
	ClientMap.Range(func(k, v interface{}) bool {
		if v == value {
			key = k.(string)

			// 跳出循环
			return false
		}

		return true
	})

	return key
}

// Delete2 is delete item form ClientMap
//
// input:
//	conn: one value
func (m SyncMap) Delete2(value interface{}) {
	m.Range(func(k, v interface{}) bool {
		if v == value {
			m.Delete(k)

			// 跳出循环
			return false
		}

		return true
	})
}

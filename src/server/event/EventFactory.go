package event

import (
	"GoWorkspace/go_line_chat/src/common/message"
	"net"
)

// EventFactory is event factory
type EventFactory struct {
	eventMap map[string]EventInterface
}

// NewEventFactory is static create EventFactory Object
func NewEventFactory() *EventFactory {
	ef := new(EventFactory)
	ef.init()
	return ef
}

func (ef *EventFactory) init() {
	// 创建map
	ef.eventMap = make(map[string]EventInterface)

	// 为map添加事件
	ef.eventMap[message.MSGSignIn] = new(EventSignIn)
}

// Parse is parse json string, to dispenser msg
func (ef *EventFactory) Parse(jsonReslut map[string]interface{}, client net.Conn) {
	// 判断json数据是否合法
	var t = jsonReslut["type"]
	if t == nil {
		return
	}

	// 分发消息给事件处理
	var ei EventInterface
	switch t {
	case message.MSGSignIn:
		ei = new(EventSignIn)
	}

	// 调用事件对象解析数据。
	if ei != nil {
		ei.Parse(jsonReslut, client)
	}
}

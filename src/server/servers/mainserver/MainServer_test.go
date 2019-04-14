// Package mainserver 主服务测试类
// name: MainServer_test.go
// 介绍: 主服务端结构体的测试文件。
package mainserver

import "testing"

func TestMainServer_Start(t *testing.T) {
	ms := new(MainServer)
	defer ms.Stop()

	ms.Start()
}

func TestStartMainServer(t *testing.T) {
	StartMainServer(8888)
}

// 测试Restart
func TestMainServer_Restart(t *testing.T) {
	ms := new(MainServer)
	defer ms.Stop()

	ms.Restart()
}

func Benchmark(b *testing.B) {
	StartMainServer(8888)
}

// ExampleStartMainServer 主服务测试类
func ExampleStartMainServer() {
	// 启动主线程
	StartMainServer(8888)
}

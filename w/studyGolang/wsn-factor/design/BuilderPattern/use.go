package main

import "fmt"

func main() {
	// 使用消息建造者进行对象创建
	msg := Builder().
		WithSrcAddr("192.168.0.1").
		WithSrcPort(1234).
		WithDestAddr("192.168.0.2").
		WithDestPort(8080).
		WithHeaderItem("contents", "application/json").
		WithBodyItem("record1").
		WithBodyItem("record2").
		Build()

	fmt.Println(msg)
}

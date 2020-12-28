package main

import (
	"bufio"
	"fmt"
	"github.com/prometheus/common/log"
	"net"
	"time"
)

var quitSemaphore chan bool

func main() {

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9078")
	if err != nil {
		log.Error("创建连接段失败")
		return
	}
	tcp, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Error("创建客户端失败")
		return
	}
	defer tcp.Close()
	//接收消息
	go onMessageRecived(tcp)
	//发送消息
	go sendMessage(tcp)
	<-quitSemaphore
}

// 发送消息
func sendMessage(conn *net.TCPConn) {
	//发送消息
	for {
		time.Sleep(1 * time.Second)
		var msg string
		fmt.Scanln(&msg)
		if msg == "quit" {
			quitSemaphore <- true
			break
		}
		//lk
		//b :=[]byte(msg +"\n")
		//处理加密
		//b ,_ := codec.Encode(msg+"\n")
		_, _ = conn.Write([]byte(msg + "\n"))
	}
}

func onMessageRecived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	for {
		//解密
		//msg,err := codec.Decode(reader)//reader.ReadString('\n')
		fmt.Println(reader.ReadString('\n'))
		//if err !=nil {
		quitSemaphore <- true
		break
		//}
	}
}

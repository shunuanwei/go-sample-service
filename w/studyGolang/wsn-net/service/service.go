package main

import (
	"bufio"
	"fmt"
	"github.com/prometheus/common/log"
	"net"
)

// 用来记录所有的客户端连接
var ConnMap map[string]*net.TCPConn

func main() {
	//初始化
	ConnMap = make(map[string]*net.TCPConn)
	//创建一个tcp的服务端
	addr, err2 := net.ResolveTCPAddr("tcp", "127.0.0.1:9078")
	if err2 != nil {
		log.Error("创建服务端 失败")
		return
	}
	//开始监听
	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Error("创建TCP 监听失败")
		return
	}
	defer listen.Close()
	for {
		tcpConn, err := listen.AcceptTCP()
		if err != nil {
			continue
		}
		fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())
		// 新连接加入 map
		ConnMap[tcpConn.RemoteAddr().String()] = tcpConn

		go tcpPipe(tcpConn)
	}

}

//处理发送过来的消息
func tcpPipe(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("disconnected : " + ipStr)
		conn.Close()
	}()
	//读取数据
	reader := bufio.NewReader(conn)
	for {
		readString, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		fmt.Println("读取到数据 ==>" + string(readString))
		//这里返回消息改为广播
		boradcastMessage(conn.RemoteAddr().String() + ":" + string(readString))
	}
}

//广播给其它
func boradcastMessage(message string) {
	//遍历所有客户端并发消息
	for _, conn := range ConnMap {
		//b,err :=codec.Encode(message)
		//if err != nil {
		//	continue
		//}
		_, _ = conn.Write([]byte(message))
	}
}

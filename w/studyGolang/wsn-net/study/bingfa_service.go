package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

//并发通讯
//每一个链接,都但单独启动一个 go程

func main() {

	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9876")

	listen, _ := net.ListenTCP("tcp", addr)
	defer listen.Close()

	for true {
		fmt.Println("服务器端 ==>>>> 等待客户端链接......")
		accept, _ := listen.Accept()

		go HandleConnect(accept)

	}
}

func HandleConnect(conn net.Conn) {
	//客户端地址
	addr := conn.RemoteAddr()
	fmt.Println(addr.String(), " ===>> 客户端链接成功!!")
	defer func(s string) {
		conn.Close()
		fmt.Printf("%s ==>> 断开链接", s)
	}(addr.String())

	//读取客户端内容
	//buf := make([]byte, 1024)

	for true {
		//read, err := conn.Read(buf)
		readString, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			if err == io.EOF {
				fmt.Println(addr.String(), "=>读取完成")
				break
			} else {
				fmt.Printf("conn.Read() err:%v\n", err)
				return
			}
		}
		//fmt.Printf("服务器读到 %s 数据:%v\n", addr.String(), string(buf[:read]))
		fmt.Printf("服务器读到 %s 数据:%v", addr.String(), readString)
		_, _ = conn.Write([]byte("已收到~~"))
	}

}

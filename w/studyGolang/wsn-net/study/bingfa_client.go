package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9876")

	conn, _ := net.DialTCP("tcp", nil, addr)

	defer conn.Close()

	//从键盘获取输入内容,发送到服务器端
	//启动一个go程,做到与接收同步
	go func() {
		for true {
			//从键盘获取 方法一:
			//str := make([]byte, 1024) 放在for循环之外 是正确的的选择
			//read, _ := os.Stdin.Read(str)
			//_, _ = conn.Write(str[:read])
			//从键盘获取  方法二:
			input := bufio.NewReader(os.Stdin)
			str, _ := input.ReadString('\n')
			_, _ = conn.Write([]byte(str))
		}
	}()

	for true {
		bytes := make([]byte, 1024)
		read, err := conn.Read(bytes)
		if err != nil {
			if err == io.EOF {
				fmt.Println("服务器退出")
				return
			} else {
				fmt.Println("err : ", err.Error())
				return
			}
		}

		fmt.Printf("客户端读到服务器回发数据: %s", bytes[:read])
	}

}

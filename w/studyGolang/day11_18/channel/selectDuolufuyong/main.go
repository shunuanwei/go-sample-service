package main

import (
	"fmt"
	"strconv"
)

/**
select   多路复用
*/
func main() {

	intChannel := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChannel <- i
	}

	strChannel := make(chan string, 5)
	for i := 0; i < 5; i++ {
		strChannel <- "hello + " + strconv.Itoa(i)
	}

	for {
		select {
		case v := <-intChannel:
			fmt.Println(v)
		case s := <-strChannel:
			fmt.Println(s)
		default:
			fmt.Println("数据获取完毕")
			return
		}
	}

}

/**
select  是随机从 channel中获取 数据
 使用 for  select   可以不关闭 channel
场景 : 从多个channel中 获取数据
*/

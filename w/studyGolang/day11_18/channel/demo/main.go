package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

//向 channel中存放数据
func putNum(intChannel chan int) {
	defer wg.Done()
	for i := 2; i < 100; i++ {
		intChannel <- i
	}
	close(intChannel)
}

//判断 质数
func paramNum(intChannel, paramChannel chan int, exitChannel chan bool) {
	defer wg.Done()
	for v := range intChannel {
		var temp = true
		for j := 2; j < v; j++ {
			if v%j == 0 {
				temp = false
				break
			}
		}
		if temp {
			paramChannel <- v
		}
	}
	exitChannel <- true

}

//打印质数
func printfParam(paramChannel chan int) {
	defer wg.Done()
	for v := range paramChannel {
		fmt.Printf("打印质数 : %v\n", v)
	}
}

func main() {
	//存放数字
	intChannel := make(chan int, 1000)
	//质数
	paramChannel := make(chan int, 1000)
	//标记一个可以关闭 质数的channel 即 不再往paramChannel 存放数据
	//标识 paramChannel 什么时间关闭
	exitChannel := make(chan bool, 16)

	wg.Add(1)
	go putNum(intChannel)

	for i := 0; i < 16; i++ {
		wg.Add(1)
		go paramNum(intChannel, paramChannel, exitChannel)
	}

	//打印质数
	wg.Add(1)
	go printfParam(paramChannel)

	//判断 可以关闭 paramChannel
	wg.Add(1)
	go func() {
		for i := 0; i < 16; i++ {
			<-exitChannel
		}
		close(paramChannel)
		wg.Done()
	}()

	//等待结束
	wg.Wait()
}

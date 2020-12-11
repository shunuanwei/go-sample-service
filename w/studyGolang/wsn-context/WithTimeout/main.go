package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	background := context.Background()

	timeoutContext, _ := context.WithTimeout(background, 5*time.Second)

	go watch(timeoutContext, "[监控一]")
	go watch(timeoutContext, "[监控二]")

	fmt.Println("现在开始等待8秒,time=", time.Now().Unix())
	time.Sleep(8 * time.Second)

	//fmt.Println("等待8秒结束,准备调用cancel()函数，发现两个子协程已经结束了，time=", time.Now().Unix())
	//cancelFunc()

}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出，停止了...")
			return
		default:
			fmt.Println(name, "goroutine监控中...")
			time.Sleep(time.Second)
		}
	}
}

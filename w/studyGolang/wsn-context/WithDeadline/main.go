package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	t := time.Now().Add(5 * time.Second)
	//cc, cancelFunc := context.WithDeadline(ctx, t)
	cc, _ := context.WithDeadline(ctx, t)

	go watch(cc, "【监控1】")
	go watch(cc, "【监控2】")
	//cancelFunc()

	time.Sleep(6 * time.Second)

}
func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "超时或者主动退出，停止了...")
			return
		default:
			fmt.Println(name, "goroutine监控中...")
			time.Sleep(time.Second)
		}
	}
}

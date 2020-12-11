package main

import "context"
import "time"
import "fmt"

func main() {
	//根context
	background := context.Background()
	ctx, cancel := context.WithCancel(background)
	go watch(ctx, "【监控1】")
	go watch(ctx, "【监控2】")
	go watch(ctx, "【监控3】")
	time.Sleep(5 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
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
	/**
	例子中启动了3个监控goroutine进行不断的监控，每一个都使用Context进行跟踪，
	当我们使用cancel函数通知取消时候，这3个 goroutine都会被结束。
	所有基于这个context或者衍生出来的子Context都会收到通知，
	这样就可以进行清理操作最终释放goroutine了
	*/
}

package main

import "fmt"

/**
通过通信实现共享内存,而不是通过共享内存实现通信
channel 先入先出

如果想要for range channel   就必须要有 close
*/

func test(n int, ch chan int) {
	for i := (n-1)*30 + 1; i <= n*30; i++ {
		if i == 1 {
			continue
		}
		temp := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				temp = false
				break
			}
		}
		if temp {
			//fmt.Println(i)
		}
	}
	ch <- n
}

func main() {

	var ch = make(chan int, 4)
	for i := 1; i <= 4; i++ {
		go test(i, ch)
	}
	/*	//for i := 1; i <= 4; i++ {
		//	ch <- i
		//}

		//close(ch)
		// for range  xunhuan
		//for v := range ch {
		//	fmt.Println(v)
		//}
	*/

	for i := 1; i <= 4; i++ {
		fmt.Println(<-ch)
	}
}

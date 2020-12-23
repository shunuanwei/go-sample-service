package main

import (
	"fmt"
	"time"
)

func main() {
	//	tt := time.Now()
	//
	//	fmt.Println(tt.Format("2006/01/02 15:04:05"))
	//
	//
	////time.Unix(时间戳,0)
	//
	//
	//	tk := time.Tick(time.Second)
	//	for i := range tk{
	//		fmt.Println(i)
	//	}

	//duration := time.Duration(10)
	//fmt.Printf("%T\n",duration)
	//fmt.Println(duration.Seconds())

	fmt.Println(time.Now().Unix())

	tk := time.NewTicker(time.Second)
	a := 0
	in := make(chan int, 1)

	go func() {
		for true {
			<-tk.C
			a++
			if a >= 5 {
				in <- 9
				//tk.Stop()
			}
			fmt.Println(a)
		}

	}()

	for true {
		<-tk.C
		if <-in > 0 {
			fmt.Println(<-in)
			fmt.Println(time.Now().Unix())
			break
		}
	}

}

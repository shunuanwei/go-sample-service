package main

import (
	"fmt"
	"sync"
	"time"
)

//协程等待

var wg sync.WaitGroup

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test : printf - ", i)
		time.Sleep(time.Millisecond * 10)
	}
	wg.Done()
}
func main() {
	wg.Add(1)
	go test()

	wg.Wait()

}

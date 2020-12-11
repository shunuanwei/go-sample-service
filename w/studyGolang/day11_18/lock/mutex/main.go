package main

import (
	"fmt"
	"sync"
	"time"
)

var mx sync.Mutex
var wg sync.WaitGroup

var count int = 0

func test() {
	defer wg.Done()
	mx.Lock()
	count++
	fmt.Println("count is ", count)
	mx.Unlock()
	time.Sleep(time.Millisecond)
}

func main() {

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go test()
		//test()
	}

	wg.Wait()

}

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var rw sync.RWMutex

func write() {
	rw.Lock()
	fmt.Println("执行写操作  <-")
	time.Sleep(time.Second * 2)
	wg.Done()
	rw.Unlock()
}

func read() {
	rw.RLock()
	fmt.Println("<-   执行读操作")
	time.Sleep(time.Second * 2)
	wg.Done()
	rw.RUnlock()
	rw.RLocker()

}

func main() {

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
}

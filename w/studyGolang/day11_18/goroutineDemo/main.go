package main

import (
	"fmt"
	"sync"
)

var wa sync.WaitGroup
var count int

func test(n int) {
	defer wa.Done()
	for i := (n-1)*30000 + 1; i <= n*30000; i++ {
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
			count++
		}
	}

}

func main() {

	for i := 1; i <= 4; i++ {
		wa.Add(1)
		go test(i)
	}
	wa.Wait()
	fmt.Println(count)

	/**
	这个还是 有问题的

	*/

}

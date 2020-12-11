package main

import "fmt"

/**
单向管道
*/

func read(ch <-chan int) {
	fmt.Println(<-ch)
}

func write(ch chan<- int) {
	ch <- 1

}
func main() {
	//ch1 := make(chan <- int,2)
	//
	//ch2 := make( <- chan int , 2)

	ch := make(chan int, 1)

	write(ch)

	read(ch)

}

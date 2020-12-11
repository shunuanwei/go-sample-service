package main

import "fmt"

func main() {
	fmt.Println("111")

	test()
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("产生异常")
		}
	}()
	//var m map[string]string
	//m["0"] = "2"

	var temp []int
	temp = append(temp, 2)
	fmt.Println(temp)
}

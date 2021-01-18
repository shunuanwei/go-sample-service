package main

import "fmt"

func main() {
	fmt.Println("111")

	test()

	fmt.Println("234234")
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("产生异常")
		}
	}()

	var temp []int
	temp = append(temp, 2)
	fmt.Println(temp)

	//没有分配内存,会报错
	var m map[string]string
	m["0"] = "2"

}

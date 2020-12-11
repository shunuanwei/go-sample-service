package calc

import "fmt"

//init 函数 优于main函数执行
func init() {
	fmt.Println("calc init")
}

func Add(x, y int) int {
	return x + y
}

func Sub(x, y int) int {
	return x - y
}

func Demo() {
	fmt.Println("hello calc")
}

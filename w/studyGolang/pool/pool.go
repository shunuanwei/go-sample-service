package main

import (
	"fmt"
	"sync"
)

var pool *sync.Pool

type Person struct {
	Name string
}

func initPool() {

	pool = &sync.Pool{
		New: func() interface{} {
			return new(Person)
		},
	}
}

func main() {
	//初始化 : Pool
	initPool()

	p := pool.Get().(*Person)

	p.Name = "张三"

	pool.Put(p)

	fmt.Println(pool.Get().(*Person).Name)

}

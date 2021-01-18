package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func main() {
	defer func() {
		if e := recover(); e != nil {

		}
	}()
	mm := new(sync.Map)

	mm.Store("1223", "3332")
	mm.Store("12233", "3333")
	mm.Store("122334", "3334")
	mm.Store(123, "3331")

	mm.Range(func(key, value interface{}) bool {
		if s, ok := key.(string); ok {
			if compare := strings.Compare(s, "123"); compare == 0 {
				mm.Store("123", "1111")
			}
		}
		return true
	})
	fmt.Println("========")

	mm.Range(func(key, value interface{}) bool {
		if s, ok := key.(string); ok {
			fmt.Println("key : " + s)
			fmt.Println("value: " + value.(string))
		} else {
			fmt.Println("key ==: " + strconv.Itoa(key.(int)))
			fmt.Println("value: " + value.(string))
		}
		return true
	})
}

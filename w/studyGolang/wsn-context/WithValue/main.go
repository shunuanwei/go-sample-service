package main

import "context"
import "fmt"

func main() {
	background := context.Background()
	ctx1 := context.WithValue(background, "01", "hello")
	ctx2 := test1(ctx1)
	test2(ctx2)
	//test2(ctx1)
}

func test1(ctx context.Context) context.Context {
	ctx02 := context.WithValue(ctx, "02", "world")
	return ctx02
}

func test2(ctx context.Context) {
	fmt.Println(ctx.Value("01"))
	fmt.Println(ctx.Value("02"))
}

package main

import (
	"context"
	"fmt"
)

func main() {

	ct := context.Background()

	fmt.Printf("%T\n", ct)
	fmt.Printf("%#v", ct)

}

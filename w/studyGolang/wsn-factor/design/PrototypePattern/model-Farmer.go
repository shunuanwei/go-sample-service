package main

import "fmt"

type Farmer struct {
}

func (f *Farmer) Work(task *string) {
	fmt.Println("Farmer process", *task)
}

type FarmerCreator struct {
}

func (c *FarmerCreator) Create() Worker {
	s := new(Farmer)
	return s
}

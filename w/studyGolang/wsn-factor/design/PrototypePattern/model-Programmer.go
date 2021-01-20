package main

import "fmt"

type Programmer struct {
}

func (p *Programmer) Work(task *string) {
	fmt.Println("Programmer process", *task)
}

type ProgrammerCreator struct {
}

func (c *ProgrammerCreator) Create() Worker {
	s := new(Programmer)
	return s
}

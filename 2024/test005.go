package main

import (
	"fmt"
)

type employee struct {
    salary  float32
}

func (e *employee) giveRaise(percent float32) {
	e.salary = e.salary * (1 + percent) 
}

func main() {
	e1 := employee{salary: 10000}
	e1.giveRaise(0.1)
	fmt.Println(e1.salary)
}
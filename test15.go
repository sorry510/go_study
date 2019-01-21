package main

import (
	"fmt"
)

// interface

type Man interface {
	name() string
	age() int
}

type Woman struct {
	height int
	id int
	city string
	fire bool
}

func (woman Woman) name() string {
	fmt.Println(woman)
	return "Helli"
}

func (woman Woman) age() int {
	return 18
}

func main() {
	var man1,man2 Man
	woman := Woman{height: 15, id: 1, city: "newyork"}
	woman.fire = true
	man1 = new(Woman)
	man2 = woman
	fmt.Println(man1.name())
	fmt.Println(man2.name())
}
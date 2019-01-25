package main

import (
	"fmt"
)

// interface

type Man interface {
	name(pre string) string
	age() int
}

type Woman struct {
	flag string
	height int
	id int
	city string
	fire bool
}

func (woman Woman) name(pre string) string {
	return woman.flag + " " + pre + " world"
}

func (woman Woman) age() int {
	return 18
}

func main() {
	man1 := new(Woman)
	fmt.Println(man1.name("hello"))

	man2 := Woman{height: 15, id: 1, city: "newyork", flag: "nv"}
	man2.fire = true
	fmt.Println(man2.name("not"))
}
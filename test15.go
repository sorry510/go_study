package main

import (
	"fmt"
)

// interface

type Man interface {
	setName(name string)
	getName(pre string) string
	setAge(age int)
	getAge() int
}

type Child struct {
	name string
	age  int
}

// go 不支持默认参数
func (self *Child) setName(name string) {
	self.name = name
}

func (self *Child) getName(pre string) string {
	return pre + self.name
}

func (self *Child) setAge(age int) {
	self.age = age
}

func (self *Child) getAge() int {
	return self.age
}

func main() {
	man1 := new(Child)
	man1.setAge(10)
	man1.setName("kelly")
	fmt.Println(*man1)
	fmt.Println(man1.getName("little "), man1.getAge())

	man2 := Child{name: "wu", age: 8}
	fmt.Println(man2.getName("big "), man2.getAge())
}

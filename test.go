package main

import "fmt"

type Is interface {
	Lost(author string) bool
}

type Book struct {
	title string
	price float64
	date string
	author string
}

func (book Book) Lost(author string) bool {
	if book.author == author {
		return true
	}
	return false
}

func main() {
	book1 := Book{"money", 10, "1900-02-12", "sorry510"}
	var l = book1.Lost("sorry510")
	fmt.Println(l)
}


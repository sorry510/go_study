package main

import (
	"fmt"
)

func main() {
	fmt.Println(f())
}

func f() (b int) {
    defer func() {
        b++
    }()
    return 1
}

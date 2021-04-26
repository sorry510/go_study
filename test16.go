package main

import (
	"fmt"
	"time"
)

// goroutine 协程
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("1")
	go say("2")
	say("3")
	// fmt.Println(1)
}

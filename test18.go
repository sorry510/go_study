package main

import (
	"fmt"
	"time"
)

// channel()

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	time.Sleep(3000 * time.Millisecond)
	c <- sum // 把 sum 发送到通道 c
}

func main() {
	fmt.Println(time.Now())
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从通道 c 中接收

	fmt.Println(x, y, x+y)
	fmt.Println(time.Now())
}

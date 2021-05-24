package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	time.Sleep(3000 * time.Millisecond)
	c <- sum // 把 sum 发送到通道 c
}

func main() {
	start := time.Now()
	c := make(chan int) // 一个无缓存的channel,可以指定第二个整型参数，对应channel的容量
	// 一个基于无缓存Channels的发送操作将导致发送者goroutine阻塞,直到另一个goroutine在相同的Channels上执行接收操作
	go sum([]int{1, 2, 3}, c)
	go sum([]int{4, 5, 6}, c)
	x, y := <-c, <-c // 从通道 c 中接收
	fmt.Println(x + y)
	secs := time.Since(start).Seconds()
	fmt.Printf("is %f seconds", secs)
}

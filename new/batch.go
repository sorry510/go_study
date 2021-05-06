package main

import (
	"fmt"
	"time"
)

// 并发
func main() {
	start := time.Now()
	ch := make(chan int)
	numbers := []int{1, 2, 3, 4, 5}
	for _, v := range numbers {
		go batch(v, ch)
	}
	sum := 0
	for range numbers {
		sum += <-ch
	}
	fmt.Println(sum)
	secs := time.Since(start).Seconds()
	fmt.Println(secs)
}

func batch(num int, ch chan<- int) {
	time.Sleep(1000 * time.Millisecond)
	ch <- num // 把 sum 发送到通道 c
}

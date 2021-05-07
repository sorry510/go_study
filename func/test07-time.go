package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)                              // 添加WaitGroup计数器
	time.AfterFunc(1*time.Second, func() { // 只运行此处异步代码，主进程会很快结束，导致无法内部代码，需用手动阻塞
		fmt.Println("1 second")
		wg.Done()
	})
	wg.Wait() // 阻塞，直到WaitGroup中的计数器为0
	// time.Sleep(4 * time.Second)
}

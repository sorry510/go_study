package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex
var m = make(map[string]int)

// 一般而言，当panic异常发生时，程序会中断运行，并立即执行在该goroutine中被延迟的函数（defer 机制）
func lookup(key string) int {
	mu.Lock()
	defer mu.Unlock() // defer语句经常被用于处理成对的操作，如打开、关闭、连接、断开连接、加锁、释放锁。通过defer机制，不论函数逻辑多复杂，都能保证在任何执行路径下，资源被释放
	return m[key]
}

func main() {
	const day = 24 * time.Hour
	fmt.Println(day, day.Seconds())
}

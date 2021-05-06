package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Unix())                                      // 当前时间戳
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))               // 当前格式化时间,记忆方法:6-1-2-3-4-5
	fmt.Println(time.Unix(1389058332, 0).Format("2006-01-02 15:04:05")) // 格式化时间戳

	start := time.Now()
	time.Sleep(3000 * time.Millisecond)
	secs := time.Since(start).Seconds()
	fmt.Println(secs)
}

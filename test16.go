package main

import (
    "time"
    "fmt"
    "math/rand"
    "strconv"
    "os"
)

// goroutine 协程

func say(s string) {
    for i := 0; i < 5; i++ {
        randTime := rand.Intn(1000)
        time.Sleep(time.Duration(randTime) * time.Millisecond)
        fmt.Println(s + strconv.Itoa(randTime))
    }
}


// 在main函数前执行
func init(){
    //以时间作为初始化种子
    rand.Seed(time.Now().UnixNano())
}

func main() {
    go say("world")
    say("hello")
    os.Exit(0)
}
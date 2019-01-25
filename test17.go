package main

import "fmt"

// 异常
// 多个panic只会捕捉最后一个

func main(){
    defer func(){
        if err := recover() ; err != nil {
            fmt.Println(err)
        }
    }()

    defer func(){
        fmt.Println("first")
        panic("three")
    }()

    defer func(){
        fmt.Println("second")
        panic("two")
    }()

    fmt.Println("third")
    panic("one")
}
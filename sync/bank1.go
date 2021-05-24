package main

import (
	"fmt"
	"time"
)

var balance int

func Deposit(amount int) {
	balance = balance + amount
}

func Balance() int {
	return balance
}

func main() {
	// 类似于在mysql没有添加事务，结果可能有3中，100，200,300
	go func() {
		Deposit(100)
	}()
	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
	}()
	time.Sleep(100 * time.Millisecond)
}

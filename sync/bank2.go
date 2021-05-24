package main

import (
	"time"
)

// Go的口头禅“不要使用共享数据来通信；使用通信来共享数据”
var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func main() {
	go Deposit(200)
	go Deposit(100)
	go teller() // start the monitor goroutine
	time.Sleep(100 * time.Millisecond)
}

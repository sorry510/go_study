package main

import "sync"

// 互斥锁
var (
	mu      sync.Mutex // guards balance
	balance int
)

func Deposit(amount int) {
	mu.Lock() // 切记对一个已经锁上的mutex来再次上锁——这会导致程序死锁，go里没有重入锁
	defer mu.Unlock()
	balance = balance + amount
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	b := balance
	return b
}

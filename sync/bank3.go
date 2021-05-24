package main

var (
	sema    = make(chan struct{}, 1) // a binary semaphore guarding balance
	balance int
)

// 手动实现 互斥锁
func Deposit(amount int) {
	sema <- struct{}{} // acquire token
	balance = balance + amount
	<-sema // release token
}

func Balance() int {
	sema <- struct{}{} // acquire token
	b := balance
	<-sema // release token
	return b
}

func main() {

}

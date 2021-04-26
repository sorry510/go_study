package main

import "fmt"

const NUM int = 3

// 数组指针

func main() {
	var a = []int{1, 2, 3}

	fmt.Println(a[1:2]) // 切片

	var ptr [NUM]*int

	for i := 0; i < NUM; i++ {
		ptr[i] = &a[i]
	}

	for i := 0; i < NUM; i++ {
		fmt.Println(*ptr[i])
	}
}

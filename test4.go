package main

import(
    "fmt"
)

// for循环
// { 不能单独起一行

func main() {
	for i := 0; i < 100; i++ {
		if i % 10 == 0 {
			fmt.Println(i)
		}
	}

    numbers := [6]int{1, 2, 3, 5}

    for i, x:= range numbers {
        fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
    }   
}
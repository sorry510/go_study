package main

import "fmt"

// range

func main() {
	nums := []int{1, 2, 3, 4, 5}
	sum := 0

	for k, v := range nums {
		sum += v
		fmt.Println(k)
	}
	fmt.Println(sum)

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for v, k := range kvs {
		fmt.Println(v, k)
	}
}

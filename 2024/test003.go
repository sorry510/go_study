package main

import (
	"fmt"
)


func main() {
	// items := []int{10, 20, 30, 40, 50}
	// // items := new([5]int){10, 20, 30, 40, 50}
	// for i, item := range items {
	// 	items[i] = item * 2
	// 	// item *= 2
	// }
	// fmt.Println(items)
	
	capitals := map[string] string {"France":"Paris", "Italy":"Rome", "Japan":"Tokyo" }
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}
	
	// a := [...]string{"a", "b", "c", "d"}
	// for i := range a {
	// 	fmt.Println("Array item", i, "is", a[i])
	// }
}
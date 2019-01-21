package main

import "fmt"

// 数组

func main(){
	// var arr1 = []float32 {1,2,3,4,5}
	var arr2 = [5]int{}
	arr2[1] = 1
	println(len(arr2), arr2[1])
	fmt.Println(arr2)
}
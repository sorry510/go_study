package main

import "fmt"

// slice
func main() {
   var numbers = make([]int,9)

   printSlice(numbers)
   printSlice(numbers[3:4])

   s :=[] int {1,2,3,4}
   printSlice(s[:2])
   printSlice(s[3:])

}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
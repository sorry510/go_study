package main

import "fmt"

var identifier []int

// slice
func main() {
    var numbers = make([]int, 9)

    printSlice(numbers)
    printSlice(numbers[3:4])

    // if(numbers[3:3] == nil){
    //     fmt.Printf("切片是空的")
    // }

    /* 同时添加多个元素 */
    numbers = append(numbers, 2, 3, 4)
    printSlice(numbers)

    /* 创建切片 numbers1 是之前切片的两倍容量*/
    numbers1 := make([]int, len(numbers), cap(numbers) * 2)

    /* 拷贝 numbers 的内容到 numbers1 */
    copy(numbers1,numbers)

    s := [] int {1, 2, 3, 4}
    printSlice(s[:2])
    printSlice(s[3:])

}

func printSlice(x []int){
    fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
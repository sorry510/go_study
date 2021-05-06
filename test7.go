package main

import (
	"fmt"
	"time"
)

type Detail struct {
	Number int
	User   string
	Title  string
	Age    time.Time
}

// 数组
var arr [10]int
var mulitArr = [3][4]int{
	{0, 1, 2, 3},   /*  第一行索引为 0 */
	{4, 5, 6, 7},   /*  第二行索引为 1 */
	{8, 9, 10, 11}, /* 第三行索引为 2 */
}

var details = []Detail{
	{Number: 1, User: "zhang", Title: "teacher", Age: time.Now()},
	{Number: 2, User: "li", Title: "docker", Age: time.Now()},
}

func main() {
	fmt.Println(arr)
	var arr1 = []float32{1, 2, 3, 4, 5}
	fmt.Println(arr1)
	var arr2 = [5]int{}
	arr2[1] = 1
	println(len(arr2), arr2[1])
	fmt.Println(arr2)
	fmt.Println(details)
}

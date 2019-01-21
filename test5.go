package main

var i, j = 100, 60

// 函数使用

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	println(max(i, j))
}

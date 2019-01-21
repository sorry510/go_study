package main

var i, j = "hello", "world"

// 函数使用多返回值

func swap(a, b string) (string, string) {
	return b, a
}

func main() {
	x, y := swap(i, j)
	println(x, y)
}

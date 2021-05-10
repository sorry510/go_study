package main

import "bytes"

var any interface{} // 空接口类型

func main() {
	any = true
	any = 12.34
	any = "hello"
	any = map[string]int{"one": 1}
	any = new(bytes.Buffer)
}

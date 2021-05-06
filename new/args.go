package main

import (
	"fmt"
	"os"
	"strings"
)

// 命令行参数可从os包的Args变量获取
func main() {
	var s, sep string

	// method1
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	// method2
	s = ""
	sep = ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	// method3
	fmt.Println(strings.Join(os.Args[1:], " "))
}

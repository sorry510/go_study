package main

import (
	"fmt"
	"strings"
)

func square(n int) int {
	return n * n
}

func sum(vales ...int) int {
	total := 0
	for _, val := range vales {
		total += val
	}
	return total
}

func main() {
	// f := square(3)
	// fmt.Println(f)

	s := strings.Map(func(r rune) rune { return r + 1 }, "hal-9000")
	fmt.Println(s)

	fmt.Println(sum(1, 2, 3, 4))
}

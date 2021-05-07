package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	res1 := Distance(p, q) // 实际上用的是包级别的函数main.Distance
	res2 := p.Distance(q)
	disFunc := p.Distance
	res3 := disFunc(q)
	fmt.Println(res1, res2, res3)

	disFunc2 := Point.Distance
	fmt.Println(disFunc2(p, q))
}

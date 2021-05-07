package main

import (
	"fmt"
	"image/color"
	"math"
	"sync"
)

type Point struct{ X, Y float64 }

type ColorPoint struct {
	Point
	Color color.RGBA
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func Lookup(key string) string {
	cache.Lock() // sync.Mutex 的方法
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

func main() {
	var cp ColorPoint
	cp.X = 1
	fmt.Println(cp.Point.X, cp.X)

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColorPoint{Color: red, Point: Point{1, 1}}
	var q = ColorPoint{Point{5, 4}, blue}
	fmt.Println(p, q)
	fmt.Println(p.Distance(q.Point)) // Point 上的方法也可以调用
}

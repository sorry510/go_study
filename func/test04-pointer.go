package main

import "fmt"

type Point struct {
	X float64
	Y float64
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	r1 := &Point{1, 2}
	r1.ScaleBy(2)
	fmt.Println(*r1)

	r2 := new(Point)
	r2.X = 1
	r2.Y = 2
	r2.ScaleBy(3)
	fmt.Println(*r2)

	r3 := Point{1, 2}
	(&r3).ScaleBy(4)
	fmt.Println(r3)

	r4 := Point{1, 2}
	r4.ScaleBy(5) // 编译器会隐式地帮我们用&r4去调用ScaleBy这个方法,这种简写方法只适用于定义好的“变量”,不能是临时变量
	fmt.Println(r4)
}

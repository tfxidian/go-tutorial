package main

import "fmt"

type Point struct {
	X, Y float64
}

func (p2 Point) isAbove(m float64) bool {
	return p2.Y > m
}

func IsAboveFunc(p Point, m float64) bool {
	return p.Y > m
}

func TranslateFunc(p *Point, dx, dy float64) {
	p.X = p.X + dx
	p.Y = p.Y + dy
}

func main() {
	p := Point{2.0, 3.0}
	fmt.Println(p.isAbove(4.0))
	//第二种方法
	fmt.Println(IsAboveFunc(p, 4.0))

	// function with结构体指针
	TranslateFunc(&p, 4, 5)
	fmt.Println(p)
}

package main

import "fmt"

type shape interface {
	getArea() float64
}

type triange struct {
	base   float64
	height float64
}
type square struct {
	sideLength float64
}

func main() {
	t := triange{base: 10, height: 10}
	s := square{sideLength: 10}

	printArea(t)
	printArea(s)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triange) getArea() float64 {
	return 0.5 * t.base * t.height
}

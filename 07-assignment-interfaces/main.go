package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

type shape interface {
	getName() string
	getArea() float64
}

func (s square) getName() string {
	return "square"
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (s triangle) getName() string {
	return "triangle"
}

func (t triangle) getArea() float64 {
	return t.base * t.height / 2
}

func printArea(s shape) {
	fmt.Printf("Area of shape '%v': %v\n", s.getName(), s.getArea())
}

func main() {
	t := triangle{base: 4.5, height: 2.3}
	s := square{sideLength: 5.0}

	printArea(t)
	printArea(s)

}

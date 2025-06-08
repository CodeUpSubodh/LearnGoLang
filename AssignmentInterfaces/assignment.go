package main

import "fmt"

type square struct {
	sideLength float64
}

type triangle struct {
	base   float64
	hieght float64
}

type shapes interface {
	area() float64
}

func (s square) area() float64 {
	area := s.sideLength * s.sideLength
	return area
}

func (t triangle) area() float64 {
	area := t.base * t.hieght * 1 / 2
	return area
}

func printArea(s shapes) {
	fmt.Println(s.area())
}

func main() {
	square := square{2.5}
	triangle := triangle{3, 4}
	printArea(square)
	printArea(triangle)

}

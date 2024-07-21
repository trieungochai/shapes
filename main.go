package main

import "fmt"

type triangle struct {
	base  float64
	heigh float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	// creating instances
	t := triangle{base: 10, heigh: 5}
	s := square{sideLength: 4}

	// calling `printArea`
	// printArea(t)/printArea(s) passes a triangle/square instance.
	// The printArea function treats t, s as a shape and calls t.getArea().
	// Go's type system recognizes that t and s satisfy the shape interface
	// because they both have the getArea method.
	printArea(t)
	printArea(s)
}

// Interface Implementation
// Both triangle and square satisfy the shape interface
// because they each implement the getArea method.
func (t triangle) getArea() float64 {
	return (t.heigh * t.base) / 2
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

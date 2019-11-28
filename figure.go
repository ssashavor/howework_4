package main

import (
	"fmt"
	"math"
)

type Figure interface {
	area() (float64, error)
	perimeter() (float64, error)
}

type Square struct {
	a float64
}
type Circle struct {
	r float64
}

func (s Square) area() (float64, error) {
	if s.a < 0 {
		return 0, fmt.Errorf("Side %f of square can't be negative(area).", s.a)
	}
	return s.a * s.a, nil
}
func (c Circle) area() (float64, error) {
	if c.r < 0 {
		return 0, fmt.Errorf("Radius %f can't be negative(area).", c.r)
	}
	return math.Pi * c.r * c.r, nil
}
func (s Square) perimeter() (float64, error) {
	if s.a < 0 {
		return 0, fmt.Errorf("Side %f of square can't be negative(perimeter).", s.a)
	}
	return s.a * 4, nil
}
func (c Circle) perimeter() (float64, error) {
	if c.r < 0 {
		return 0, fmt.Errorf("Radius %f can't be negative(perimeter).", c.r)
	}
	return 2 * math.Pi * c.r, nil
}

func main() {
	var s Figure = Square{2}
	var c Figure = Circle{-5}
	var sO Figure = Square{-8}
	var cO Figure = Circle{5}

	fmt.Println(sO.area())
	fmt.Println(s.area())
	fmt.Println(cO.area())
	fmt.Println(c.perimeter())
	fmt.Println(s.perimeter())
	fmt.Println(cO.perimeter())
}

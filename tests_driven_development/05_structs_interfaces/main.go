package main

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base  float64
	Height float64
}

// methods
func (r Rectangle) Area() float64 {
	res := r.Width * r.Height
	return res
}
func (c Circle) Area() float64 {
	res := math.Pi * c.Radius * c.Radius
	return res
}
func (t Triangle) Area() float64 {
	res := (t.Height * t.Base) / 2
	return res
}

func Perimeter(rectangle Rectangle) float64 {
	res := 2 * (rectangle.Width + rectangle.Height)
	return res
}

// func Area(rectangle Rectangle) float64 {
// 	res := rectangle.Width * rectangle.Height
// 	return res
// }

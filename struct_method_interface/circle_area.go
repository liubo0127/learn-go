package main

import "math"

type Rectangle struct {
	height float64
	weight float64
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.height + r.weight)
}

func (r *Rectangle) Area() float64 {
	return r.height * r.weight
}

type Circle struct {
	Radius float64
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Shape interface {
	Perimeter() float64
	Area() float64
}

func main() {
}
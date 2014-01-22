package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

type MultiShape struct {
	shapes []Shape // It's possible to have an interface as a field.
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func main() {
	c := &Circle{0, 0, 5}
	fmt.Println(c.area())

	r := &Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())

	m := &MultiShape{}
	m.shapes = append(m.shapes, c)
	m.shapes = append(m.shapes, r)
	fmt.Println(m.area())

	// The area method of each A, B and C are different entirely.
	fmt.Println(m.shapes[0].area())
	fmt.Println(m.shapes[1].area())
}

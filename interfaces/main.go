package main

import (
	"fmt"
	"math"
)

type Shaper interface {
	Shape() string
}

type Area interface {
	Area() float64
}

type Rectangle struct {
	len, br float64
}

func (r *Rectangle) Area() float64 { return r.len * r.br }
func (r *Rectangle) Shape() string { return "rectangle" }

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 { return math.Pi * math.Pow(c.radius, 2) }
func (c *Circle) Shape() string { return "circcle" }

func main() {
	var s1, s2 Shaper = &Rectangle{2, 4}, &Circle{3}
	var a1, a2  = s1.(Area), s2.(Area)
	fmt.Println(a1.Area(), a2.Area())
}

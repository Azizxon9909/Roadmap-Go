package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle1 struct {
	radius float64
}

type rect1 struct {
	width  float64
	height float64
}

func (r rect1) area() float64 {
	return r.width * r.height
}

func (c circle1) area() float64 {
	return math.Pi * c.radius * c.radius
}

func getArea(s shape) float64 {
	return s.area()
}

func main1() {
	c1 := circle1{4.5}
	r1 := rect1{5, 7}
	shapes := []shape{c1, r1}
	for _, v := range shapes {
		fmt.Println(getArea(v))
	}	
}

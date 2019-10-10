package main

import (
	"math"
	"fmt"
)

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	var c1 circle
	c1.radius = 5
	fmt.Println(c1.area())
}
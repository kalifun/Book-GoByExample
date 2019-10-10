package main

import (
	"fmt"
	"math"
)

type geometry interface {
	aera() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) aera() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (x circle) aera() float64 {
	return math.Pi * x.radius * x.radius
}

func (x circle) perim() float64 {
	return 2 * math.Pi * x.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.aera())
	fmt.Println(g.perim())
}

func main() {
	s := rect{width: 2, height: 3}

	d := circle{radius: 4}
	measure(s)
	measure(d)
}

package examples

import (
	"fmt"
	"math"
)

type object interface {
	name() string
}

type geometry interface {
	object
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2 * (r.width + r.height)
}

func (r rect) name() string {
	return "Rectangular"
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) name() string {
	return "Circle"
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("area: ", g.area())
	fmt.Println("perim: ", g.perim())
	fmt.Println("name: ", g.name())
}

func GoInterfaces() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	measure(r)
	measure(c)
}

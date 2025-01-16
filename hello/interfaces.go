package hello

import "fmt"

type geometry interface {
	area() float64
	perim() float64
}
type reactangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r reactangle) area() float64 {
	return r.width * r.height
}

func (r reactangle) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * 3.14 * c.radius
}
func Measure(g geometry) {
	println(g)
	println(g.area())
	println(g.perim())
}
func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func InterfaceFunc() {
	r := reactangle{width: 3, height: 4}
	c := circle{radius: 5}

	Measure(r)
	Measure(c)
	detectCircle(c)
}

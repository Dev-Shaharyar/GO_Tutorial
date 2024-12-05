package interfaces

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

//Interfaces  -> code clean , redable whenever a data come from a 3rd
//party API we dont need to check the what are the values they have

// {
// 	name : "string"
// 	id :int
// }
// {
// 	name :"string"
// } will be decide by the interfaces

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}
type square struct {
	side float64
}

// rectangle
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// for the circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// for the square
func (s square) area() float64 {
	return s.side * s.side
}

func (s square) perim() float64 {
	return 4 * s.side
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func Interfaces() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	s := square{side: 5}

	measure(r)
	measure(c)
	measure(s)
}

package main

import (
	"fmt"
	"math"
)

// interface in go is used to override the inherit in OO

type Cicle struct {
	name   string
	Radius float64
}

type Rectangle struct {
	name          string
	Width, Height float64
}

// the interface is a custom type that
// is used to specify a set of one or more method signatures
type Shape interface {
	area() float64
	changeName(newName string)
}

func main() {
	circle := Cicle{
		Radius: 4.6,
	}

	rec := Rectangle{
		Width:  8.9,
		Height: 5.0,
	}

	// slice of structs
	// with interface we can do a slice of objects
	shapes := []Shape{&circle, &rec}

	// only to go through the shape
	for _, shape := range shapes {
		fmt.Println(shape.area())
	}

	fmt.Println("eaeaea")
}

// overriding the method from the interface Shape
func (c *Cicle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// overriding the method from the interface Shape
func (r *Rectangle) area() float64 {
	return r.Width * r.Height
}

func (c *Cicle) changeName(value string) {
	c.name = value
}

func (r *Rectangle) changeName(value string) {
	r.name = value
}

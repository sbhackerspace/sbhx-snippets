// Steve Phillips / elimisteve
// 2013.10.13

package main

import (
	"log"
)

// Example from http://pragprog.com/articles/erlang

type Shape interface {
	Area() float64
}

// Rectangle
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Square
type Square struct {
	Width float64
}

func (sq Square) Area() float64 {
	return sq.Width * sq.Width
}

// Circle
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func main() {
	shapes := []Shape{
		Rectangle{10, 5},
		Square{7.3},
		Circle{1.4},
	}

	if rect, ok := shapes[0].(Rectangle); ok {
		log.Printf("%#v is a Rectangle\n", rect)
	}

	if rect, ok := shapes[0].(Shape); ok {
		log.Printf("%+v is a Shape\n", rect)
	}

	for _, s := range shapes {
		log.Printf("Area of %T:\t%v\n", s, s.Area())
	}
}

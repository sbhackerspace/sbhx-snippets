// Steve Phillips / elimisteve
// 2010.12.15

package main

import (
	"fmt"
	"math"
)

type Point struct { x, y float64 }

func (p Point) Abs(a, b float64) float64 {
	p.x, p.y = a, b
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func setValues(pnt *Point, x float64, y float64) {
	(*pnt).x, (*pnt).y = x, y
}

func main() {
	 // Struct
	var p Point
	var xcor, ycor float64 = 3, 4

	 // Custom method Abs
	fmt.Printf("(%.1f, %.1f) distance from origin: %.1f\n\n",
		xcor, ycor, p.Abs(xcor, ycor))

	 // Pointers
	fmt.Printf("Before setValues\n")
	fmt.Printf("p.x = %.1f, p.y = %.1f\n\n", p.x, p.y)

	setValues(&p, xcor, ycor)
	
	fmt.Printf("After setValues\n")
	fmt.Printf("p.x = %.1f, p.y = %.1f\n", p.x, p.y)
}
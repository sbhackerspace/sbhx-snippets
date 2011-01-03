// Steve Phillips / elimisteve / fraktil
// 2010.12.14

package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Point struct { x, y float64 }

func Distance(p Point) float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func main() {
	var p Point
	var erx, ery os.Error
	p.x, erx = strconv.Atof64(flag.Arg(0))
	p.y, ery = strconv.Atof64(flag.Arg(1))

	if erx != nil {
		fmt.Printf("Error: %s\n", erx)
	} else if ery != nil {
		fmt.Printf("Error: %s\n", ery)
	} else {
		fmt.Printf("p.x is %.1f\n", p.x)
		fmt.Printf("p.y is %.1f\n", p.y)
		fmt.Printf("Distance from origin: %.1f\n", Distance(p))
	}
}
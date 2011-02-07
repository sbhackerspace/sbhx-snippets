// Steve Phillips / elimisteve
// 2011.02.06

package main

import (
	"fmt"
)

type MyFloat float64

func (n MyFloat) Square() float64 {
	x := float64(n)
	return x * x
}

func main() {
	var n MyFloat = 4
	fmt.Printf("%.2f\n", n.Square())
}
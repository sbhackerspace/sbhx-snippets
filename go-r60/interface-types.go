// Steve Phillips / elimisteve
// 2012.02.06

package main

import (
	"fmt"
)

func main() {
	// `holdsAnything` can be the type of anything that implements the
	// interface `interface{}`
	var holdsAnything interface{}

	holdsAnything = 5
	fmt.Printf("%v\n", holdsAnything)

	holdsAnything = "This is a string"
	fmt.Printf("%v\n", holdsAnything)

	holdsAnything = true
	fmt.Printf("%v\n", holdsAnything)
}

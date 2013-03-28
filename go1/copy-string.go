// Steve Phillips / elimisteve
// 2013.03.26

package main

import (
	"fmt"
)

func main() {
	orig := "This is a string"
	newer := orig
	newer += " as well"
	fmt.Printf("%v\n", orig)
	fmt.Printf("%v\n", newer)
}
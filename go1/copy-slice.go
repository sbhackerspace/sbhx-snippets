// Steve Phillips / elimisteve
// 2013.03.26

package main

import (
	"fmt"
)

func main() {
	orig := []byte("This is a string")
	newer := orig
	newer = append(newer, []byte(" as well")...)
	fmt.Printf("%s\n", orig)
	fmt.Printf("%s\n", newer)
}
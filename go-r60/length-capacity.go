// Steve Phillips / elimisteve
// 2011.02.27

package main

import (
	"fmt"
)

func main() {
	s := make([]int, 3, 5)
	fmt.Printf("s == %v\n", s)
	fmt.Printf("s[3:] == %v\n", s[3:])
	s = s[:cap(s)]
	fmt.Printf("Length increased from 3 to 5\n")
	s[4] = 1
	fmt.Printf("s[3:] == %v\n", s[3:])
}
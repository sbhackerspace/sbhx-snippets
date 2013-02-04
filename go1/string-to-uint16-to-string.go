// Steve Phillips / elimisteve
// 2013.02.03

package main

import (
	"fmt"
	"unicode/utf16"
)

func main() {
	// Start with a string
	s1 := "Degrees °; Copyright ©; Registered ®"
	fmt.Printf("%s\n", s1)

	// Convert to []uint16
	uints := utf16.Encode([]rune(s1))
	fmt.Printf("%c\n", uints)

	// Convert back to string
	slice := []rune{}
	for _, c := range uints {
		slice = append(slice, rune(c))
	}
	s2 := string(slice)
	fmt.Printf("%s\n", s2)
}
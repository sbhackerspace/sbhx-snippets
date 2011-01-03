// Steve Phillips / elimisteve / fraktil
// 2010.12.14

package main

import (
	"fmt"
	"flag"
)

func main() {
	 // for, range
	for index, arg := range flag.Args() {
		println("flag.Arg", index, "is", flag.Arg(index))
		println("flag.Arg", index, "is", arg)
	}
	fmt.Printf("\n")

	 // String concatenation
	s1 := ""
	fmt.Printf("Concatenated command line args:\n")
	for i := 0; i < flag.NArg(); i++ {
		s1 += flag.Arg(i)
	}
	fmt.Printf("%s\n\n", s1)

	 // Slices
	s2 := ""
	fmt.Printf("Some chars from each arg:\n")
	for i := 0; i < flag.NArg(); i++ {
		s2 += flag.Arg(i)[1:3] + " "
	}
	fmt.Printf("%s\n\n", s2)

	s3 := s1
	fmt.Printf("The characters of s3:\n")
	for i := 0; i < len(s3); i++ {
		fmt.Printf("%c\n", s3[i]) // %c, not %s
	}
}
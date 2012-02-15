// Steve Phillips / elimisteve
// 2012.02.06

package main

import (
	"fmt"
)

type Person struct {
	FirstName, LastName string
	Age int
}

func main() {
	p := Person{FirstName: "Steve", LastName: "Phillips", Age: 26}
	fmt.Println(p) // Alternatively, fmt.Printf("%v\n", p)
}
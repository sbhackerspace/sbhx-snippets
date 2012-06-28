// Steve Phillips / elimisteve
// 2012.06.27

package main

import (
	"fmt"
)

// Person
type Person struct {
	First, Last, Gender string
	Age, Weight int
}

func (p Person) FullName() string {
	return p.First + " " + p.Last
}

// Human, which is a Person with one extra field
type Human struct {
	Person
	Middle string
}

// NOTE: Comment this method out, then run program again
func (h Human) FullName() string {
	return h.First + " " + h.Middle + " " + h.Last
}

func main() {
	// Person
	p := Person{First: "Steven", Last: "P", Age: 26}
	fmt.Printf("%v\n", p.FullName())
	// Human
	h := Human{
		Person: Person{First: "Ashley", Last: "P", Age: 17},
		Middle: "Ann",
	}

	// This works; method names don't conflict. Comment out
	// Human.FullName() method to see Person.FullName() called instead.
	fmt.Printf("%v\n", h.FullName())
}
// Steve Phillips / elimisteve
// 2012.02.06

package main

import (
	"fmt"
	"strings"
)

type Printable interface {
	Print() string
}

type IntSlice []int

func (slice IntSlice) Print() string {
	sliceStr := fmt.Sprintf("%v", slice)
	stringSlice := strings.Split(sliceStr, " ")
	return strings.Join(stringSlice, ", ")
}

type MyInt int

func (i MyInt) Print() string {
	return fmt.Sprintf("%v", i)
}


func main() {
	implementMsg := "%v is of type %T and has value %v\n"

	slice := IntSlice{1, 2, 3}
	number := MyInt(5)
	stuff := []interface{}{slice, number}

	for _, item := range stuff {
		// Detect the type of item
		switch t := item.(type) {
		// If item is of type MyInt or IntSlice...
		case MyInt, IntSlice:
			// ...and implements the Printable interface...
			if printMe, ok := item.(Printable); ok {
				// ...print it
				fmt.Printf(implementMsg, item, item, printMe.Print())
			}
		}
	}
}
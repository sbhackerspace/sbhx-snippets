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
	implementMsg := "%v indeed implements %v and has value %v\n"

	slice := IntSlice{1, 2, 3}
	number := MyInt(5)
	stuff := []interface{}{slice, number}

	for _, item := range stuff {
		// Next line won't work because empty interfaces don't have
		// the .Print() method
		// fmt.Printf(implementMsg, item, "Printable", item.Print())
		if printMe, ok := item.(Printable); ok {
			fmt.Printf(implementMsg, item, "Printable", printMe.Print())
		}
	}
}
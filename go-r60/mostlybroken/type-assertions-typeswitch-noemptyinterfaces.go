// Steve Phillips / elimisteve
// 2012.02.06

package main

import (
	"fmt"
)

type IntSlice []int

func main() {
	//
	// Nevermind: can't be done! Type switches are on interface{} values
	//
	
	slice := IntSlice{1, 2, 3}
	num   := 5
	msg   := "%v is indeed a(n) %T with value %v\n"

	switch t := slice.(type) {
	default:
		fmt.Printf(msg, slice, slice, slice)
	}

	switch t := num.(type) {
	default:
		fmt.Printf(msg, num, num, num)
	}
}
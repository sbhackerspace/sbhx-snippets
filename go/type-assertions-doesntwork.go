// Steve Phillips / elimisteve
// 2012.02.06

package main

import (
	"fmt"
)

type IntSlice []int

func main() {
	//
	// The following is completely wrong. See
	// type-assertions-interfaces.go or type-assertions-typeswitch.go
	//

	slice := IntSlice{1, 2, 3}
	if sl, ok := slice.(IntSlice); ok {
		fmt.Printf("%v is indeed an IntSlice with value %v\n", slice, ok)
	}

	number := 5
	if num, ok := number.(int); ok {
		fmt.Printf("%v is indeed an int with value %v\n", number, num)
	}

}
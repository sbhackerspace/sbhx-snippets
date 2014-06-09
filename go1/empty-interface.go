// Steve Phillips / elimisteve
// 2014.06.08

package main

import (
	"fmt"
)

func main() {
	lst := []interface{}{"string", 5, 3.14159265}
	for i, val := range lst {
		fmt.Printf("%v %v\n", i, val)
	}
}

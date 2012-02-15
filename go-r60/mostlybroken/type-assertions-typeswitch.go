// Steve Phillips / elimisteve
// 2012.02.06

package main

import (
	"fmt"
)

type IntSlice []int

func main() {
	slice  := IntSlice{1, 2, 3}
	number := 5
	stuff  := []interface{}{slice, number}
	msg    := "%v is indeed a(n) %v with value %v\n"

	for _, item := range stuff {
		switch t := item.(type) {
		case int:
			fmt.Printf(msg, item, "int", item)
		case IntSlice:
			fmt.Printf(msg, item, "IntSlice", item)
		}
	}
}
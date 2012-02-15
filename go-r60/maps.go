// Steve Phillips / elimisteve / fraktil
// 2010.12.17

package main

import (
	"fmt"
)

var words = map[string] string { "2":"hundred", "3":"thousand" }

func main() {
	value3, ok := words["3"]
	if ok {
		fmt.Printf("%s\n", value3)
	}

	value4, ok := words["4"]
	if ok {
		fmt.Printf("%s\n", value4)
	}
}
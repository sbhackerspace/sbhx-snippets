// Steve Phillips / elimisteve
// 2012.02.06

// From areyoufuckingcoding.me/2012/02/06/netchan-is-no-more-shall-be-chanio/

package main

import (
	"fmt"
	"os"
)

func main() {
	x := interface{}("some string")
	// `string` is a type, not an interface!!... right?
	if str, ok := x.(string); ok {
		fmt.Println(str)
	} else {
		fmt.Fprintf(os.Stderr, "%v is not a string!", x)
	}
}

// Steve Phillips / elimisteve
// 2011.09.07

package main

import (
    "fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
// Steve Phillips / elimisteve
// 2011.03.26

package main

import (
	"fmt"
)

func main() {
    // Slices: k, v == index, value at index
    for k, v := range []string{"zero", "one", "two"} {
        fmt.Printf("%v %v\n", k, v)
    }
    println()

    // Maps: k, v == key, value
    factorials := map[int]int{0: 1, 1: 1, 2: 2, 3: 6, 4: 24}
    for k, v := range factorials {
        fmt.Printf("%v! = %v\n", k, v)
    }
    println()

    // Strings: k, v == index, char at index
    my_string := "Hello, world"
    for k, v := range my_string {
        fmt.Printf("%s is at %d\n", string(v), k)
    }

}
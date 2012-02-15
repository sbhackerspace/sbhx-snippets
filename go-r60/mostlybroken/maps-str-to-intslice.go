// Steve Phillips / elimisteve
// 2011.06.01

package main

import (
    "fmt"
)

func f() map[string][]int {
    // Ugly... Good thing only toy examples look like this.
    return map[string][]int {
        "a": []int{1,2,3,4,5},
        "b": []int{6,7,8,9,10},
    }
}

func main() {
    fmt.Printf("%v\n", f())
}
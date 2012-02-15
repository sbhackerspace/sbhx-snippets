// Steve Phillips / elimisteve
// 2011.08.01

package main

import (
    "fmt"
)

type intSlice []int

func (slice intSlice) AppendTill(max int) intSlice {
    last := slice[len(slice)-1]
    // Base Case
    if last == max {
        return intSlice(slice)
    }
    // Recursive Call
    return append(slice, last+1).AppendTill(max)
}

func main() {
    slice := intSlice( []int{1,2,3} )
    fmt.Printf("Before: %v\n", slice)
    slice = slice.AppendTill(10)
    fmt.Printf("After:  %v\n", slice)
}
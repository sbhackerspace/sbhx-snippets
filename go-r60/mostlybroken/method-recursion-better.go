// Steve Phillips / elimisteve
// 2011.08.01

package main

import (
    "fmt"
)

type intSlice []int

func (slice intSlice) AppendTill(max int) intSlice {
    last := slice[len(slice)-1]
    switch {
    case last == max:
        return intSlice(slice)
    case last < max:
        return append(slice, last+1).AppendTill(max)
    case last > max:
        return append(slice, last-1).AppendTill(max)
    }
    return slice
}

func main() {
    slice := intSlice( []int{1} )
    fmt.Printf("Start:     %v\n", slice)
    slice = slice.AppendTill(10)
    fmt.Printf("Up to 10:  %v\n", slice)
    slice = slice.AppendTill(1)
    fmt.Printf("Back down: %v\n", slice)
}
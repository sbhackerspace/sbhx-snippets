// Steve Phillips / elimisteve
// 2011.08.01

package main

import (
    "fmt"
)

type intSlice []int

func (sl intSlice) AppendNum() intSlice {
    slice := []int(sl)
    return intSlice( append(slice, slice[len(slice)-1]+1) )
}

func main() {
    slice := intSlice( []int{1,2,3} )
    fmt.Printf("Before: %v\n", slice)
    for slice[len(slice)-1] < 10 {
        slice = slice.AppendNum()
    }
    fmt.Printf("After:  %v\n", slice)
}
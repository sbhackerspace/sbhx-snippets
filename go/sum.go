// Steve Phillips / elimisteve
// 2011.04.17

package main

import (
    "fmt"
)

func sum(nums []int) (total int) {
    for _, n := range nums {
        total += n
    }
    return
}

type MySlice []int

func (s MySlice) Sum() (total int) {
    slice := []int(s)
    for _, n := range slice {
        total += n
    }
    return
}

func main() {
    intSlice := []int{1,2,3}
    fmt.Printf("%d\n", sum(intSlice))
    fmt.Printf("%d\n", MySlice(intSlice).Sum())
}

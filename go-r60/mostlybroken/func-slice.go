// Steve Phillips / elimisteve
// 2011.05.29

package main

import (
    "fmt"
)

//add := func(a, b int) int { return a + b }
func add(nums []int) (sum int) {
    for _, num := range nums {
        sum += num
    }
    return
}

//subtract := func(a, b int) int { return a - b }
func subtract(nums []int) (difference int) {
    if len(nums) == 0 {
        return 0
    }
    difference = nums[0]
    for _, num := range nums[1:] {
        difference -= num
    }
    return
}

func mult(nums []int) (product int) {
    if len(nums) == 0 {
        return 0
    }
    product = 1
    for _, num := range nums {
        product *= num
    }
    return
}

type intSlice []int

func (ints intSlice) Reverse() []int {
    if len(ints) == 0 {
        return ints
    }
    sl := []int(ints)
    reversed := []int{}
    end := len(sl)-1
    for ndx, _ := range sl {
        reversed = append(reversed, sl[end-ndx])
    }
    return reversed
}

func main() {
    for _, f := range []func([]int) int {add, subtract, mult} {
        fmt.Printf("%v\n", f( intSlice([]int{1, 2, 3, 6}).Reverse() ))
    }
}
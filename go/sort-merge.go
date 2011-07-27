// Steve Phillips / elimisteve
// 2011.07.26

package main

import (
    "fmt"
)

func MergeSort(numbers []int) []int {
    length := len(numbers)
    if length < 2 {
        return numbers
    }
    left   := MergeSort( numbers[:length/2] )
    right  := MergeSort( numbers[length/2:] )
    // left and right are now sorted. Time to merge.

    sorted := []int{}
    l, r := 0, 0  // left and right indices/"indexes"
    for {
        if left[l] < right[r] {
            sorted = append(sorted, left[l])
            l++
            if l == len(left) {
                // if done with left, append rest of right's elements
                sorted = append(sorted, right[r:]...)
                break
            }
        } else {
            sorted = append(sorted, right[r])
            r++
            if r == len(right) {
                // if done with right, append rest of left's elements
                sorted = append(sorted, left[l:]...)
                break
            }
        }
    }
    return sorted
}

func main() {
    nums := []int{12, 70, -43, 100, 34, 111, 94, 25, 0, 7, -2, 113, -40, 40}
    fmt.Printf("Before: %v\n", nums)
    fmt.Printf("After:  %v\n", MergeSort(nums))
}
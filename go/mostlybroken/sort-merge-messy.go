// Steve Phillips / elimisteve
// 2011.07.26

package main

import (
    "fmt"
)

// // merge takes two int slices and 
// func sorted(numbers []int) []int {
    
// }

func MergeSort(numbers []int) []int {
    length := len(numbers)
    if length < 2 {
        return numbers
    }
    left, right := numbers[:length/2], numbers[length/2:]
    //left, right  = sorted(left), sorted(right)
    //lp, rp := &left[0], &right[0]
    result := []int{}

    //for l, r := 0, 0; len(result) < len(left)+len(right);  {
//    for l, r := 0, 0; len(result) < len(left)+len(right) && l < len(left) && r < len(right);  {
    for l, r := 0, 0; l + r < len(result);  {
        if left[l] < right[r] {
            result = append(result, left[l])
            l++
        } else {
            result = append(result, right[r])
            r++
        }
    }

    //return MergeSort( MergeSort(left), MergeSort(right)... )
    return result
}

func main() {
    nums := []int{12, 70, -43, 100}//, 34, 94, 111, 25, 0, 2, 7, -17, 40}
    fmt.Printf("%v\n", MergeSort(nums))
}
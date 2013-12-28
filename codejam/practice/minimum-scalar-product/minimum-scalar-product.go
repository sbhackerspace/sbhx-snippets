// Steve Phillips / elimisteve
// 2013.12.28

package main

import (
	"fmt"

	jam "github.com/elimisteve/go.jam"
)

func main() {
	nCases := jam.GetInt()
	for i := 0; i < nCases; i++ {
		nDim := jam.GetInt()
		// Populate v1
		v1 := make([]int, nDim)
		for i := 0; i < nDim; i++ {
			v1[i] = jam.GetInt()
		}
		// Populate v2
		v2 := make([]int, nDim)
		for i := 0; i < nDim; i++ {
			v2[i] = jam.GetInt()
		}
		v1 = quicksort(v1, true)
		v2 = quicksort(v2, false)
		fmt.Printf("Case #%d: %v\n", i+1, scalarProduct(v1, v2))
	}
}

func quicksort(nums []int, ascend bool) []int {
	if len(nums) < 2 {
		return nums
	}
	left, right := split(nums[0], nums[1:], ascend)
	sorted := append(quicksort(left, ascend), nums[0])
	return append(sorted, quicksort(right, ascend)...)
}

func split(head int, rest []int, ascend bool) (left, right []int) {
	for i := 0; i < len(rest); i++ {
		if ascend {
			// Small to big
			if rest[i] <= head {
				left = append(left, rest[i])
			} else {
				right = append(right, rest[i])
			}
		} else {
			// Big to small
			if rest[i] >= head {
				left = append([]int{rest[i]}, left...)
			} else {
				right = append([]int{rest[i]}, right...)
			}
		}
	}
	return
}

func scalarProduct(v1, v2 []int) (prod int) {
	// Assumes len(v1) == len(v2)
	for i := 0; i < len(v1); i++ {
		prod += v1[i] * v2[i]
	}
	return
}

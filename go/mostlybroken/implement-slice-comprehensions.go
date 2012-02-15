// Steve Phillips / elimisteve
// 2011.11.11

package main

import (
    "fmt"
)

func between(start, end int) (numbers []int) {
    for i := start; i <= end; i++ {
		numbers = append(numbers, i)
	}
	return
}

func sliceSum(nums []int) (total int) {
	for _, num := range nums {
		total += num
	}
	return
}

func intFilter(claim func(i int) bool, nums []int) (passed []int) {
	for _, num := range nums {
		if claim(num) {
			passed = append(passed, num)
		}
	}
	return
}

func main() {
	euler1 := func(i int) bool { return i%3 == 0 || i%5 == 0 }
	nums := intFilter(euler1, between(1,999))
	fmt.Printf("Total: %d\n", sliceSum(nums))
}
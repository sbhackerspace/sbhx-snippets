// Steve Phillips / elimisteve
// 2013.12.27

package main

import (
	"fmt"
	jam "github.com/elimisteve/go.jam"
)

func main() {
	nCases := jam.GetInt()
	for i := 0; i < nCases; i++ {
		credit := jam.GetInt()
		nItems := jam.GetInt()
		prices := jam.GetInts(nItems)
		positions := findSum(credit, prices)
		fmt.Printf("Case #%d: %d %d\n", i+1, positions[0], positions[1])
	}
}

func findSum(credit int, prices []int) [2]int {
	// O(n^2); improve this
	for i := 0; i < len(prices); i++ {
		for j := 0; j < len(prices); j++ {
			if i == j {
				continue
			}
			if prices[i]+prices[j] == credit {
				return [2]int{i, j}
			}
		}
	}
	// Not found; this shouldn't happen
	return [2]int{-1, -1}
}

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
		nFlavors := jam.GetInt()
		nCustomers := jam.GetInt()
		// For each customer, some number of "X Y" pairs, where X is
		// the flavor and Y is 1 for malted, 0 for unmalted
		likes := make([][][2]int, nCustomers)
		var T int // Number of pairs "X Y"
		for i := 0; i < nCustomers; i++ {
			T = jam.GetInt()
			likes[i] = jam.GetPairs(T)
		}

		output := "IMPOSSIBLE"
		toPrep, err := calculateBatches(nFlavors, likes)
		if err == nil {
			output = jam.FmtInts(toPrep)
		}
		fmt.Printf("Case #%d: %v\n", i+1, output)
	}
}

func calculateBatches(nFlavors int, likes [][][2]int) ([]int, error) {
	// TODO
	return make([]int, nFlavors), nil
}

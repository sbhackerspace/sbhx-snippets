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
		v1 = jam.QuicksortInts(v1, true)
		v2 = jam.QuicksortInts(v2, false)
		fmt.Printf("Case #%d: %v\n", i+1, scalarProduct(v1, v2))
	}
}

func scalarProduct(v1, v2 []int) (prod int) {
	// Assumes len(v1) == len(v2)
	for i := 0; i < len(v1); i++ {
		prod += v1[i] * v2[i]
	}
	return
}

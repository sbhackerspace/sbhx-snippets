// Steve Phillips
// 2010.12.13
// OCUK Stage 1

package main

import (
	"fmt"
)

func main() {

	//MAX_NUM := 100
	MAX_NUM := 1500000

	longest_chain, longest_n := 0, 0

	for num := 2; num < MAX_NUM; num++ {
		number := num
		chain_length := 0

		for number != 1 {
			if number & 1 == 0 {
				number >>= 1
			} else if number & 1 == 1 {
				number += (number << 1) + 1
			}
			chain_length += 1

			if chain_length > longest_chain {
				longest_chain = chain_length
				longest_n = num
			}
		}
	}

	fmt.Printf("n = %d generates chain length %d", longest_n, longest_chain)
}
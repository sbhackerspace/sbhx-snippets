// Steve Phillips / elimisteve
// 2012.02.02

package main

import (
	"fmt"
)

func main() {
	const NUM_LOOPS = 5
	signal := make(chan int)

	// Send numbers 0 through NUM_LOOPS-1 to channel `signal`
	go func() {
		for i := 0; i < NUM_LOOPS; i++ {
			signal <- i+1
		}
	}()

	for i := 0; i < NUM_LOOPS; i++ {
		// This loop pulls all NUM_LOOPS-many numbers from the `signal`
		// channel then does something with them
		num := <-signal
		switch {
		case num == 1:
			fmt.Printf("==1: Got %v\n", num)
		default:
			// Run when no other conditions are satisfied. Whether
			// it's first, last, or in the middle doesn't matter
			fmt.Printf("default: %v\n", num)
		case num < 4:
			// Numbers that are both < 3 and < 4 get printed here
			// because the < 4 condition occurs before < 3
			fmt.Printf("< 4: Got %v\n", num)
		case num < 3:
			// Never runs, since all numbers < 3 are also < 4, and the
			// < 4 condition occurs first
			fmt.Printf("< 3: Got %v\n", num)
		}
	}
	// `i` is undefined here because Go's variables are block-scoped
}

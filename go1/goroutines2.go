// Steve Phillips / elimisteve
// 2013.01.03

package main

import "fmt"

// intDoubler doubles the given int, then sends it through the given channel
func intDoubler(ch chan int, n int) {
	ch <- n*2
}

func main() {
	// Make a channel of ch
	ch := make(chan int)
	answer := make(chan string)

	// Spawn 3 goroutines (basically threads) to process data in background
	go intDoubler(ch, 10)
	go intDoubler(ch, 20)
	go func(a, b int) { ch <- a+b }(30, 40) // Take 2 ints; Write sum to `ch`

	// Create anonymous function on the fly, launch as goroutine!
	go func() {
		// Save the 3 values passed through the channel as x, y, and z
		x, y, z := <-ch, <-ch, <-ch
		// Create answer; write to `answer` channel
		answer <- fmt.Sprintf("%d + %d + %d = %d", x, y, z, x+y+z)
	}()

	// Print answer from channel
	fmt.Printf("%s\n", <-answer)
}

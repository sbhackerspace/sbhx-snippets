// Steve Phillips / elimisteve
// 2011.06.11

package main

import (
    "fmt"
    "io/ioutil"
    "time"
)

// intDoubler doubles then sends the given int through the given channel
func intDoubler(ch chan int, n int) {
    ch <- n*2
}

func intTripler(ch chan int, n int) {
    ch <- n*3
}

func main() {
    // Make a channel of ints
    ch := make(chan int)

    // Spawn 3 goroutines (basically threads) to process data in background
    go intDoubler(ch, 10)
    go intTripler(ch, 20)
    go func(c chan int, a, b int) { c <- a+b }(ch, 30, 40)

    // Create anonymous function on the fly, launch as goroutine!
    go func(c chan int) {
        // Save the 3 values passed through the channel as x, y, and z
        x, y, z := <-c, <-c, <-c
        // Create string containing the answer; Write it to /tmp/answer
        str := fmt.Sprintf("%d + %d + %d = %d\n", x, y, z, x+y+z)
        ioutil.WriteFile("/tmp/answer", []byte(str), 0644)
    }(ch)

    // Simulation of doing other stuff while the above is computed
    time.Sleep(2e9) // Sleep for 2 seconds == 2 x 10^9 nanoseconds

    fmt.Printf("Find the calculated answer in /tmp/answer\n")
}

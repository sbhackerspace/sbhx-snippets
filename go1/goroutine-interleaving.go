// Steve Phillips / elimisteve
// 2014.06.09

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Rough Go equivalent of concurrent C++ program explained at
// https://www.youtube.com/watch?v=80ifzK3b8QQ, but with artificial
// pauses to simulate the slowness of threads

const times = 10

func main() {
	rand.Seed(time.Now().UnixNano())

	done := make(chan bool)
	for i := 0; i < times; i++ {
		go func(num int) {
			time.Sleep(time.Duration(rand.Intn(5e8)))
			fmt.Printf("Hello from goroutine %d!\n", num)
			done <- true
		}(i)
	}
	time.Sleep(time.Duration(rand.Intn(5e8)))
	fmt.Printf("Hello from main!\n")

	for i := 0; i < times; i++ {
		<-done
	}
}

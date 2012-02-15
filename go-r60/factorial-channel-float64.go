package main

import (
//	"flag"
//	"fmt"
//	"strconv"
)

func numgen(ch chan <-float64, fromuser float64) {
	var i float64
	for i = 0; i < fromuser; i++ {
		ch <- i+1
	}
}

func factorial(n float64) float64 {
	var total float64 = 1
	var i float64

	ch := make(chan float64)
	go numgen(ch, n)
	for i = 0; i < n; i++ {
		total *= <-ch
	}
	return total
}

func main() {
	var num1 float64 = 170
	for i := 0; i < 1000; i++ {
		factorial(num1)
	}
}
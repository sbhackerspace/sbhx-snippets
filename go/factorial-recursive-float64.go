package main

import (
//	"flag"
//	"fmt"
//	"strconv"
)

func factorial(n float64) float64 {
	if n==0 || n==1 {
		return 1
	} else {
		return n * factorial(n-1)
	}
	return 1
}

func main() {
	var num1 float64 = 170
	for i := 0; i < 1000; i++ {
		factorial(num1)
	}
	// num1, _ := strconv.Atof64(flag.Arg(0))
	// fmt.Printf("%.0f\n", factorial(num1))
}
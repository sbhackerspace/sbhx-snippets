package main

import (
	"fmt"
	"flag"
	"strconv"
)

func main() {
	num1, _ := strconv.Atoi(flag.Arg(1))
	num2, _ := strconv.Atoi(flag.Arg(2))
	fmt.Printf("%d\n", num1 * num2)
	//factorial()
}
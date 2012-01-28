// Steve Phillips / elimisteve
// 2012.01.27

package main

import (
	"fmt"
	"log"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
		fmt.Printf("\nRecovered just fine! See?\n")
	}()

	intSlice := []int{0}
	fmt.Printf("%v\n", intSlice[1]) // Out of range => panic

	fmt.Printf("This won't print; program panicked on previous line\n")
}

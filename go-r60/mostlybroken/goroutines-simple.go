// Steve Phillips / elimisteve
// 2011.08.12

package main

import (
    "fmt"
)

func main() {
	done := make(chan bool)

	for i := 0; i < 10; i++ {
		go func(j int) {
			fmt.Println(j)
			done <- true
		}(i)
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}
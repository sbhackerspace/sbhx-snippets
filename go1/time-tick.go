// Steve Phillips / elimisteve
// 2013.01.04

package main

import (
	"fmt"
	"time"
)

func main() {

	tick := time.Tick(1 * time.Second)
	for {
		<-tick
		fmt.Printf("%d Printing after every tick\n", time.Now().Unix())
	}
}
// Steve Phillips / elimisteve
// 2013.06.16

package main

import (
	"log"
)

func main() {
	msgs := make(chan bool)

	go incrementor(msgs)

	for {
		msgs <- true
	}
}

func incrementor(msgs chan bool) {
	inc := 0
	for {
		<-msgs
		inc += 1
		if inc % 1e7 == 0 {
			log.Printf("%d\n", inc)
		}
	}
}
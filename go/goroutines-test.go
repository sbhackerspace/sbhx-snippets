// Steve Phillips / elimisteve
// 2011.08.24

package main

import (
    "fmt"
	"time"
)

func main() {
	done := make(chan bool)
	for _, num := range []int64{0,1,2,3,4} {
		go func(num int64) {
			fmt.Printf("%v\n", num)
			time.Sleep(1e9 * num)
			fmt.Printf("Done with loop %d\n", num)
			done <- true
		}(num)
	}
	for _, _ = range []int64{0,1,2,3,4} {
		<-done
	}
}
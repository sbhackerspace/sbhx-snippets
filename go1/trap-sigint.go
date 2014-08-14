// Steve Phillips / elimisteve
// 2014.06.19

package main

import (
	"fmt"
	"os"
	"os/signal"
	// "runtime"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go func() {
		fmt.Printf("Got this signal: %v\n", <-c)
		// runtime.Goexit()
		os.Exit(1)

		// for sig := range c {
		// 	fmt.Printf("Got this signal: %v\n", sig)
		// }
	}()
	select {}
}

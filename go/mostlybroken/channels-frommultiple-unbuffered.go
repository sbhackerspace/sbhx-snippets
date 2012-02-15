// Steve Phillips / elimisteve
// 2012.01.08

package main

func main() {
	ch1 := make(chan string, 2)
	ch2 := make(chan string, 1)

	ch1 <- "Hello, "
	ch1 <- "!"
	ch2 <- "world"

	println(<-ch1 + <-ch2 + <-ch1)
}
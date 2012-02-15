// Steve Phillips / elimisteve
// 2012.01.08

package main

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "Hello, "
		ch1 <- "!"
	}()
	go func()  {
		ch2 <- "world"
	}()

	println(<-ch1 + <-ch2 + <-ch1)
}
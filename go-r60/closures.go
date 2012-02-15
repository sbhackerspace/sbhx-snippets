// Code from http://golang.org/doc/go_faq.html#closures_and_goroutines
// Added to sbhx-snippets on 2012.01.26

package main

func main() {
	done := make(chan bool)

	values := []string{ "a", "b", "c" }
	// (From http://golang.org/doc/go_faq.html#closures_and_goroutines)
	for _, v := range values {
		go func() {
			println(v)
			done <- true
		}()
	}
	for _ = range values {
		<-done 
	}

	println()

	//
	// Better; also from http://golang.org/doc/go_faq.html#closures_and_goroutines
	//
	for _, v := range values {
		go func(s string) {
			println(s)
			done <- true
		}(v)
	}
	for _ = range values {
		<-done 
	}
}

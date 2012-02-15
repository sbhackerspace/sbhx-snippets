// Steve Phillips / elimisteve
// 2012.01.01

package main

func magicallyReturnTwo() (one int) {
	// increments `one` _after_ the returned `1` is evaluated!
	// See http://golang.org/doc/go_spec.html#Defer_statements
	defer func() { one++ }()

	return 1
}

func main() {
	println(magicallyReturnTwo())
}

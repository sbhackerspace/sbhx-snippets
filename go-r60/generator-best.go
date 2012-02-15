// Steve Phillips / elimisteve
// 2012.01.17

package main

func bestGenerator(optional ...int) func() int {
	start := 0
	if len(optional) > 0 {
		start = optional[0]
	}
	return func() int {
		start++
		return start
	}
}

func main() {
	// No parameter passed in; starts at 0
	g1 := bestGenerator()
	for i := 0; i < 10; i++ {
		print(g1(), " ")
	}
	println()

	// One parameter passed in; starts there instead
	g2 := bestGenerator(27)
	for i := 0; i < 10; i++ {
		print(g2(), " ")
	}
	println()
}
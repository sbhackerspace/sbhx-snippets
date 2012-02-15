// Steve Phillips / elimisteve
// 2012.01.17

package main

func betterGenerator(start int) func() int {
	return func() int {
		start++
		return start
	}
}

func main() {
	g := betterGenerator(5)
	for i := 0; i < 10; i++ {
		print(g(), " ")
	}
	println()
}
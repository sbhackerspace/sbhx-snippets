// Steve Phillips / elimisteve / fraktil
// 2010.12.14

package main

func return_multiple(x, y int) (int, int) {
	return x+2, y+3
}

func main() {
	a, b := return_multiple(1, 2)
	println(a, b)
}
package main

func generator() func() int {
	number := 0
	return func() int {
		number++
		return number
	}
}

func main() {
	g := generator()
	for i := 0; i < 10; i++ {
		print(g(), " ")
	}
	println()
}

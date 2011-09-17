// Steve Phillips / elimisteve
// 2011.09.07

package main

func main() {
	s := ""
	defer println(s)                 // Doesn't work
	//defer func() { println(s) }()  // Works
	s += "Seems like the output shouldn't be blank, but it is!"
}

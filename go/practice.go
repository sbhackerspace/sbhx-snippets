package main

import (
//	"os"
	"fmt"
	"flag"
)

func main() {
	string := "Hello, world"
    fmt.Printf("%s%s", string, "\n")
	println(string)

    //flag.Parse()  // What is this for?

	for i := 0; i < flag.NArg(); i++ {
		fmt.Printf("%s\n", flag.Arg(i))
	}
	//return 0
}
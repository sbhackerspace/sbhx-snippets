// Steve Phillips / elimisteve
// 2012.10.21

package main

import (
	"fmt"
	"io"
	"os"
)

func StringWriter(w io.Writer, name string) {
	// These 3 are equivalent
	fmt.Fprintf(w, "Hello, %v-person!\n", name)
	io.WriteString(w, `Hello, ` + name + "-person!\n")
	w.Write([]byte(fmt.Sprintf("Hello, %v-person!\n", name)))
}

func main() {
	StringWriter(os.Stdin, "dude")
}

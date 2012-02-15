// Steve Phillips / elimisteve
// 2011.02.06

package main

import (
    "fmt"
)

type MyInt int

func (n MyInt) Square() int {
    return int(n * n)
}

func main() {
    n := 4
    fmt.Printf("%v\n", MyInt(n).Square())
	
}
// Steve Phillips / elimisteve
// 2011.02.06

package main

import (
    "fmt"
)

type int int64

func (n int) Square() float64 {
    x := float64(n)
    return x * x
}

func main() {
    var n int = 4  // "int" refers to my custom type!
    fmt.Printf("%.2f\n", n.Square())
}

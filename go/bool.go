// Steve Phillips / elimisteve
// 2011.03.27

package main

import (
    "fmt"
)

func main() {
    var x float32 = 5

    if bool(x) {
        fmt.Printf("True\n")
    } else {
        fmt.Printf("False\n")
    }
}
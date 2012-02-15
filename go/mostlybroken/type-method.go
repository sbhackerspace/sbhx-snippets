// Steve Phillips / elimisteve
// 2011.05.13

package main

import (
    "fmt"
)

func main() {
    fmt.Printf("%v\n", four.(float32))
}

func four() int {return 4}
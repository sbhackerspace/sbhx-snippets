// Steve Phillips / elimisteve
// 2011.07.26

package main

import (
    "fmt"
)

func main() {
    items := []interface{}{0.1, 42, uint(5), float64(64), "string", 'x'}
    //fmt.Printf("%v\n", items...)  // Why doesn't this work?
    for _, v := range items {
        fmt.Printf("%v ", v)
    }
}
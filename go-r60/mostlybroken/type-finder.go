// Steve Phillips / elimisteve
// 2011.06.14

package main

import (
    "fmt"
)

func typeFinder(values ...interface{}) {
    for _, v := range values {
        switch t := v.(type) {
        default:
            fmt.Printf("%-15v is of type %T\n", v, t)
        }
    }
    return
}

func main() {
    typeFinder(5, 3.4, "hello", make(chan bool), true, nil)
}
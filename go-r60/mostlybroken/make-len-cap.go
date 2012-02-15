// Steve Phillips / elimisteve
// 2011.05.11

package main

import (
    "fmt"
)

func main() {
    type Greeting struct {
        From  string
        To    string
        Body  string
    }
    greetings := make([]Greeting, 0, 10)
    fmt.Printf("len: %v\n", len(greetings))
    fmt.Printf("cap: %v\n", cap(greetings))
}
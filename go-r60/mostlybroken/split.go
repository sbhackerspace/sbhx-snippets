// Steve Phillips / elimisteve
// 2011.05.27

package main

import (
    "fmt"
    "strings"
)

func main() {
    // Split on spaces
    sentence := "These are some words"
    words := strings.Split(sentence, " ", -1)
    for ndx, word := range words {
        fmt.Printf("word #%d is %s\n", ndx+1, word)
    }
    fmt.Printf("\n")

    // Split on commas
    sentence = "Here,are,different,words"
    words = strings.Split(sentence, ",", -1)
    for ndx, word := range words {
        fmt.Printf("word #%d is %s\n", ndx+1, word)
    }
}
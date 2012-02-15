// Steve Phillips / elimisteve
// 2011.05.31

package main

import (
    "crypto/md5"
    "fmt"
)

func main() {
    words := []string{"crap", "string", "secret"}
    hashString := "5ebe2294ecd0e0f08eab7690d2a6ee69"
    for ndx, word := range words {
        h := md5.New()
        h.Write([]byte(word))
        hash := fmt.Sprintf("%x", h.Sum())
        if hash == hashString {
            fmt.Printf("Success! Original string: %v\n", word)
            return
        } else {
            fmt.Printf("Attempt #%d failed: %s\n", ndx+1, hash)
        }
    }

}
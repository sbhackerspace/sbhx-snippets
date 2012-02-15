// Steve Phillips / elimisteve
// 2011.05.27

package main

import (
    "fmt"
    "json"
)

type Message struct {
    Original string
}

func main() {
    b := []byte(`{"Original": "password"}`)
    var m Message
    if err := json.Unmarshal(b, &m); err != nil {
        var f interface{}
        er := json.Unmarshal(b, &f)
        if er != nil {
            fmt.Printf("Something baaad happened...\n")
        }
    }
    fmt.Printf("%s\n", b)
}
// Steve Phillips / elimisteve
// 2011.03.26

package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    result, err := ioutil.ReadFile("example_file.txt")
    if err != nil {
        fmt.Printf("%v\n", err)
    } else {
        fmt.Printf("File contents\n")
        fmt.Printf("=============\n")
        fmt.Printf("%v\n", string(result))
    }
}
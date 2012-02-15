// Steve Phillips / elimisteve
// 2011.06.11

package main

import (
    "fmt"
    "io/ioutil"
    //"time"
)

func createFiles() {
    ioutil.WriteFile("/tmp/answer1", []byte("string 1\n"), 0666)
    ioutil.WriteFile("/tmp/answer2", []byte("string 2\n"), 0666)
}

func main() {
    go createFiles()
    //time.Sleep(1e9)
    fmt.Printf("Done.\n")
}

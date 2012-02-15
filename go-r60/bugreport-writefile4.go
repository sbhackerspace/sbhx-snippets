package main

import (
    "fmt"
    "io/ioutil"
)

func createFiles(ch chan bool) {
    ioutil.WriteFile("/tmp/answer1", []byte("string 1\n"), 0666)
    ioutil.WriteFile("/tmp/answer2", []byte("string 2\n"), 0666)
    ch <- true
}

func main() {
    done := make(chan bool)
    go createFiles(done)
    <-done
    fmt.Printf("Done.\n")
}
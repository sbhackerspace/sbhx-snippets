// Steve Phillips / elimisteve
// 2011.06.11

package main

import (
    "io/ioutil"
)

func writeTwoFiles() {
    ioutil.WriteFile("/tmp/answer1", []byte("This is /tmp/answer1\n"), 0644)
    ioutil.WriteFile("/tmp/answer2", []byte("This is /tmp/answer2\n"), 0644)
}

func main() {
    go writeTwoFiles()
}
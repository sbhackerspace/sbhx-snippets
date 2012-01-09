// Steve Phillips / elimisteve
// 2011.09.07

package main

import (
	"io/ioutil"
)

func main() {
	s := ""
	defer func() { ioutil.WriteFile("deferWrite.txt", []byte(s), 0644) }()
	s += "Output shouldn't be blank!\n"
	println("See ./deferWrite.txt")
}
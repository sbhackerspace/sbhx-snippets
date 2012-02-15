// Steve Phillips / elimisteve
// 2011.09.07

package main

import (
	"io/ioutil"
	"time"
)

func main() {
	s := ""
	defer ioutil.WriteFile("deferWrite.txt", []byte(s), 0644)
	s += "Output shouldn't be blank!"
	time.Sleep(2e9)
}
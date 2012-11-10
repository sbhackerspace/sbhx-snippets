// Steve Phillips / elimisteve
// 2012.11.10

package main

import (
	"bytes"
	"fmt"
	"github.com/elimisteve/fun"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

var incrementMe = 3

func main() {
	linePrefix := `var incrementMe = `

	thisFile := "/home/steve/sbhx/sbhx-snippets/go1/self-incrementer.go"
	// thisFile := os.Args[0] + ".go"
	contents, err := ioutil.ReadFile(thisFile)
	fun.MaybeFatalAt("ioutil.ReadFile", err)

	// Find the line where `incrementMe` is defined
	getNum := regexp.MustCompile(linePrefix + `\d+`)
	oldLine := getNum.FindAllString(string(contents), -1)[0]
	// for _, digits := range getNum.FindAllString(string(contents), -1) {
	// 	fmt.Printf("'%v'\n", digits)
	// }

	// Parse out the int to be incremented (the current value of `incrementMe`)
	intStr := oldLine[len(linePrefix):]
	fmt.Printf("Old value: %v\n", intStr)

	// Convert to int and increment
	theInt, err := strconv.Atoi(intStr)
	fun.MaybeFatalAt("Atoi", err)
	theInt++
	fmt.Printf("New value: %v\n", theInt)

	// Replace old value of `incrementMe` with new value
	newContents := bytes.Replace(contents, []byte(oldLine),
		[]byte(linePrefix + strconv.Itoa(theInt)), 1)

	// Open thisFile and set its contents to newContents
	file, err := os.OpenFile(thisFile, os.O_WRONLY, 0666)
	fun.MaybeFatalAt("OpenFile", err)

	_, err = file.Write(newContents)
	fun.MaybeFatalAt("file.Write", err)
}

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

const (
	LINE_PREFIX       = `var incrementMe = `
	INCREMENT_COMMAND = `go build self-incrementer.go; ./self-incrementer`
)

var incrementMe = 0

func main() {
	thisFile := os.Args[0] + ".go"
	contents, err := ioutil.ReadFile(thisFile)
	if err != nil {
		fmt.Printf("Run this command several times: \n\n    %s\n\n",
			INCREMENT_COMMAND)
		return
	}

	// Find the line where `incrementMe` is defined
	getNum := regexp.MustCompile(LINE_PREFIX + `\d+`)
	oldLine := getNum.FindAllString(string(contents), -1)[0]

	// Parse out the int to be incremented (the current value of `incrementMe`)
	intStr := oldLine[len(LINE_PREFIX):]

	// Convert to int and increment
	theInt, err := strconv.Atoi(intStr)
	fun.MaybeFatalAt("Atoi", err)
	theInt++
	fmt.Printf("%v\n", theInt)

	// Replace old value of `incrementMe` with new value
	newContents := bytes.Replace(contents, []byte(oldLine),
		[]byte(LINE_PREFIX + strconv.Itoa(theInt)), 1)

	// Open thisFile and set its contents to newContents
	file, err := os.OpenFile(thisFile, os.O_WRONLY, 0666)
	fun.MaybeFatalAt("OpenFile", err)

	_, err = file.Write(newContents)
	fun.MaybeFatalAt("file.Write", err)
}

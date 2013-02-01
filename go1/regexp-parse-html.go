// Steve Phillips / elimisteve
// Started  2011.02.06
// Improved 2013.02.01

package main

import (
	"fmt"
	"regexp"
)

func main() {
	html := `<div><p>Some comment</p></div> <div><p>Another one</p></div>`
	pat := "(?:<p>)(.*?)(?:</p>)"
	getComment := regexp.MustCompile(pat)
	for _, c := range getComment.FindAllStringSubmatch(html, -1) {
		fmt.Printf("%v\n", c[1])
	}
}

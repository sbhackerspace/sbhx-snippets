// Steve Phillips / elimisteve
// 2013.02.03

package main

import (
	"fmt"
)

func main() {
	msg := "°©®"
	// Print the same thing, but in hex
	msg += "\xc2\xb0"
	msg += "\xc2\xa9"
	msg += "\xc2\xae"
	// Again print the same thing, but using small-u Unicode code points
	msg += string([]byte{'\u00c2', '\u00b0'})
	msg += string([]byte{'\u00c2', '\u00a9'})
	msg += string([]byte{'\u00c2', '\u00ae'})
	// _Again_ print the same thing, but using BIG-u Unicode code points
	msg += string([]byte{'\U000000c2', '\U000000b0'})
	msg += string([]byte{'\U000000c2', '\U000000a9'})
	msg += string([]byte{'\U000000c2', '\U000000ae'})
	fmt.Printf(msg)
}
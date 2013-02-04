// Steve Phillips / elimisteve
// 2013.02.03

package main

import (
	"fmt"
)

func main() {
	msg := "°"
	msg += "©"
	msg += "®"
	// Hex strings
	msg += "\xc2\xb0"
	msg += "\xc2\xa9"
	msg += "\xc2\xae"
	// Small-u strings
	msg += "\u00b0"
	msg += "\u00a9"
	msg += "\u00ae"
	// Big-u strings
	msg += "\U000000b0"
	msg += "\U000000a9"
	msg += "\U000000ae"
	// Small-u Unicode code points as byte slices
	msg += string([]byte{'\u00c2', '\u00b0'})
	msg += string([]byte{'\u00c2', '\u00a9'})
	msg += string([]byte{'\u00c2', '\u00ae'})
	// Big-u Unicode code points
	msg += string([]byte{'\U000000c2', '\U000000b0'})
	msg += string([]byte{'\U000000c2', '\U000000a9'})
	msg += string([]byte{'\U000000c2', '\U000000ae'})
	fmt.Printf(msg)
}
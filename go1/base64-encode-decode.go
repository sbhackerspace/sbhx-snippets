// Steve Phillips / elimisteve
// 2013.03.03

package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	text := "Hello world?"
	fmt.Printf("Original text:  %s\n", text)

	// Encode
	enc := make([]byte, base64.StdEncoding.EncodedLen(len(text)))
	base64.StdEncoding.Encode(enc, []byte(text))
	fmt.Printf("Encoded result: %s\n", enc)

	// Decode
	b, err := base64.StdEncoding.DecodeString(string(enc))
	if err != nil {
		fmt.Errorf("Error %v", err)
	}
	fmt.Printf("Decoded result: %s\n", b)
}
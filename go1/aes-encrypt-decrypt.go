// Steve Phillips / elimisteve
// 2013.03.03

package main

import (
	"bytes"
	"crypto/aes"
	"fmt"
	"github.com/elimisteve/fun"
	"github.com/thecloakproject/utils"
	"github.com/thecloakproject/utils/crypt"
	"strings"
)

var PASSPHRASE = []byte("0123456789abcdef")

func main() {
	msg := []byte("This is 18+ chars!")
	fmt.Printf("msg ==    %s\n", msg)

	// Encrypt
	encBlock, err := aes.NewCipher(PASSPHRASE)
	fun.MaybeFatalAt("aes.NewCipher", err)

	// See https://github.com/thecloakproject/utils/blob/master/crypt/aes.go
	cipher, err := crypt.AESEncryptBytes(encBlock, msg)
	fun.MaybeFatalAt("AESEncryptBytes", err)

	fmt.Printf("cipher == %v\n", cipher)

	// Decrypt
	decBlock, err := aes.NewCipher(PASSPHRASE)
	fun.MaybeFatalAt("aes.NewCipher", err)

	// See https://github.com/thecloakproject/utils/blob/master/crypt/aes.go
	plain, err := crypt.AESDecryptBytes(decBlock, cipher)
	fun.MaybeFatalAt("AESDecryptBytes", err)

	fmt.Printf("plain ==  %s\n", plain)
	msgPadded := utils.PadBytes(msg, decBlock.BlockSize())

	// Check for equality
	fmt.Printf("\nThey match? %v!\n", bytes.Equal(msgPadded, plain))

	// Check for equality in other ways
	msgUnpadded := strings.TrimSpace(string(msgPadded))
	match := (msgUnpadded == string(plain))
	fmt.Printf("\nDo their trimmed versions match? %v!\n", match)
	if match {
		fmt.Printf("They both equal '%s'\n", msgUnpadded)
	}

	// Here's how to remove those ugly trailing nulls
	fmt.Printf("Cleanest-looking version: '%s'\n",
		strings.TrimRight(string(plain), "\x00"))
}

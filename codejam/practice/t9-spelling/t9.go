// Steve Phillips / elimisteve
// 2013.12.28

package main

import (
	"bufio"
	"fmt"
	"os"

	jam "github.com/elimisteve/go.jam"
)

func main() {
	nCases := jam.GetInt()
	r := bufio.NewReader(os.Stdin)
	for i := 0; i < nCases; i++ {
		var sequence string
		var c byte
		var err error
		for {
			c, err = r.ReadByte()
			if c == '\n' || err != nil {
				break
			}
			miniSeq := char2seq[c]
			if len(sequence) > 0 && sequence[len(sequence)-1] == miniSeq[0] {
				sequence += " "
			}
			sequence += miniSeq
		}
		fmt.Printf("Case #%d: %v\n", i+1, sequence)
	}
}

var char2seq = map[byte]string{
	' ': "0",
	'a': "2",
	'b': "22",
	'c': "222",
	'd': "3",
	'e': "33",
	'f': "333",
	'g': "4",
	'h': "44",
	'i': "444",
	'j': "5",
	'k': "55",
	'l': "555",
	'm': "6",
	'n': "66",
	'o': "666",
	'p': "7",
	'q': "77",
	'r': "777",
	's': "7777",
	't': "8",
	'u': "88",
	'v': "888",
	'w': "9",
	'x': "99",
	'y': "999",
	'z': "9999",
}

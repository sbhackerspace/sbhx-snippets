// Steve Phillips / elimisteve
// 2013.12.28

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	jam "github.com/elimisteve/go.jam"
)

func main() {
	nCases := jam.GetInt()
	r := bufio.NewReader(os.Stdin)
	for i := 0; i < nCases; i++ {
		line, _ := r.ReadString('\n')
		words := strings.Split(line[:len(line)-1], " ")
		var reversed []string
		for _, word := range words {
			reversed = append([]string{word}, reversed...)
		}
		fmt.Printf("Case #%d: %v\n", i+1, strings.Join(reversed, " "))
	}
}

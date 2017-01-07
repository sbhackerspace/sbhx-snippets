// Steve Phillips / elimisteve
// 2017.01.07

package main

import (
	"encoding/base64"
	"io"
	"os"
)

func main() {
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)

	buf := make([]byte, 0, 1e7)

	for {
		n, err := os.Stdin.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic("Error reading: " + err.Error())
		}
		if _, err = encoder.Write(buf[:n]); err != nil {
			panic("Error writing:" + err.Error())
		}
	}
}

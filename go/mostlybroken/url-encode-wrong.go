// Steve Phillips / elimisteve
// 2011.12.17

package main

import (
	"encoding/base64"
    "fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func checkError(err os.Error) {
	if err != nil {
		log.Fatalf("%s", err)
	}
}

func main() {
	encodedURL := "http%3A%2F%2Fvideo.madthumbscdn.com%2Fvideos%2F1000%2F1000874.flv%3Fsr%3D1200%26int%3D512000b%26nvb%3D20111217071227%26nva%3D20111217091227"
	// Bullshit...
	if m := len(encodedURL) % 4; m != 0 {
		encodedURL += strings.Repeat("=", 4-m)
		// fmt.Printf("Added %d ='s for padding\n", 4-m)
	}

	buf := strings.NewReader(encodedURL)
	r := base64.NewDecoder(base64.URLEncoding, buf)
	decodedURL, err := ioutil.ReadAll(r)
	checkError(err)
	fmt.Printf("%v\n", string(decodedURL))
}
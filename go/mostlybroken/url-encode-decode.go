// Steve Phillips / elimisteve
// 2011.12.17

package main

import (
	// "encoding/base64"
    "fmt"
	"log"
	"os"
	"url"
)

func checkError(err os.Error) {
	if err != nil {
		log.Fatalf("%s", err)
	}
}

func main() {
	// encodedURL := "http%3A%2F%2Fvideo.madthumbscdn.com%2Fvideos%2F1000%2F1000874.flv%3Fsr%3D1200%26int%3D512000b%26nvb%3D20111217071227%26nva%3D20111217091227"
	// encodedURL := "https%3A%2F%2Fwww.google.com%2Fsearch%3Fq=hello,+world"
	encodedURL := "https%3A%2F%2Fwww.google.com%2Fsearch%253Fq%3Dhello%2C%2Bworld"
	// encodedURL := "https://www.google.com/search?sourceid=chrome&ie=UTF-8&q=galaxy+nexus+google+wallet"

	// Original, encoded URL
	fmt.Printf("encodedURL   == %v\n", encodedURL)

	// Decode
	decodedURL, err := url.Parse(encodedURL)
	checkError(err)
	fmt.Printf("decodedURL   == %v\n", decodedURL.Path)
	fmt.Printf("decodedURL   == %v\n", decodedURL.String())

	// Re-encode
	reEncodedURL := decodedURL.EncodedPath()
	fmt.Printf("reEncodedURL == %v\n", reEncodedURL)

	// Re-decode
	reDecodedURL, err := url.Parse(reEncodedURL)
	checkError(err)
	fmt.Printf("reDecodedURL == %v\n", reDecodedURL)
}
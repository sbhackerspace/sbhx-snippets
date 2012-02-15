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
	encodedURL := "http%3A%2F%2Fvideo.madthumbscdn.com%2Fvideos%2F1000%2F1000874.flv%3Fsr%3D1200%26int%3D512000b%26nvb%3D20111217071227%26nva%3D20111217091227"
	myUrl, err := url.Parse(encodedURL)
	checkError(err)
	fmt.Printf("%s\n", myUrl.String())
}
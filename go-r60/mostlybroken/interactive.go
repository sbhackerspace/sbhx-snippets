// Steve Phillips / elimisteve
// 2012.01.09

package main

import (
	"bufio"
    "fmt"
	"http"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

func checkError(err os.Error) {
	if err != nil {
		log.Fatalf("%s", err)
	}
}

func randomQuote() string {
	// Simulates slow connection; shows off `defer`, goroutine asynchronicity
	defer func() { time.Sleep(2e9) }()

	// Scrape site
	url := "http://subfusion.net/cgi-bin/quote.pl?quote=cookie&number=1"
	req, err := http.Get(url)
	checkError(err)

	// Create regex
	re, err := regexp.Compile("<body><br><br><b><hr><br>(.*)<br><br><hr><br>")
	checkError(err)

	// Read HTML
	html, err := ioutil.ReadAll(req.Body)
	req.Body.Close()
	checkError(err)

	// Parse out quote
	quote := re.FindString(string(html))
	quote = strings.Replace(quote, "<body><br><br><b><hr><br>", "", -1)
	quote = strings.Replace(quote, "<br><br><hr><br>", "", -1)

	return quote
}

func main() {
	signal := make(chan string)
	//var input byte

	// Listen for user input, send to second goroutine
	go func() {
		for {
			r := bufio.NewReader(os.Stdin)
			input, err := r.ReadByte()
			checkError(err)
			//fmt.Scanf("%c", &input)
			switch input {
			case ' ':
				signal <- "NEWLINE"
			case 'h':
				fmt.Printf("Press 's' to stop\n")
				fmt.Printf("Press ' ' for newlines\n")
			case 'r':
				//fmt.Printf("Random quote:\n")
				fmt.Printf(randomQuote())
			case 's':
				signal <- "STOP"
			// default:
			// 	//fmt.Printf("continuing...\n")
			// 	continue
			}
		}
	}()

	// Use user input to trigger events
	go func() {
		for  {
			switch <-signal {
			case "STOP":
				fmt.Printf("Stopping...\n")
				os.Exit(0)
			case "NEWLINE":
				fmt.Printf("\n\n\n")
			}
		}
	}()

	// Main loop
    fmt.Printf("Press 'h' for help\n\n")
	for {
		fmt.Printf("%v seconds since the epoch\n", time.Seconds())
		time.Sleep(1e9)
	}
}
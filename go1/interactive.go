// Steve Phillips / elimisteve
// 2012.01.09

package main

import (
	"bufio"
    "fmt"
	"net/http"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

func checkError(err error) {
	if err != nil {
		log.Fatalf("%s", err)
	}
}

func randomQuote() string {
	// Simulates slow connection; shows off `defer`, goroutine asynchronicity
	// defer func() { time.Sleep(2e9) }()

	// Scrape site
	fmt.Printf("Scraping site...\n")
	url := "http://subfusion.net/cgi-bin/quote.pl?quote=cookie&number=1"
	req, err := http.Get(url)
	checkError(err)

	// Create regex
	re, err := regexp.Compile("<br>\n(.*)\n<br>")
	// re, err := regexp.Compile("<body><br><br><b><hr><br>(.*)<br><br><hr><br>")
	checkError(err)

	// Read HTML
	html, err := ioutil.ReadAll(req.Body)
	req.Body.Close()
	checkError(err)

	// Parse out quote
	fmt.Printf("Parsing out quote...\n")
	fmt.Printf("Original: %v\n", string(html))
	quote := re.FindString(string(html))
	// quote = strings.Replace(quote, "<body><br><br><b><hr><br>", "", -1)
	// quote = strings.Replace(quote, "<br><br><hr><br>", "", -1)
	quote = strings.Replace(quote, "<br>", "", -1)

	return "Final: " + quote + "\n"
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
			case 'r':
				fmt.Printf("Random quote:\n")
				fmt.Printf(randomQuote())
			case 's':
				signal <- "STOP"
			case 'h':
				fmt.Printf("Press ' ' for newlines\n")
				fmt.Printf("Press 'r' for random quote\n")
				fmt.Printf("Press 's' to stop\n")
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
		// fmt.Printf("It is now %v\n", time.Now())
		time.Sleep(1e9)
	}
}
// Steve Phillips / elimisteve
// 2011.01.16

package main

import (
	"fmt"
	"http"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

// Save scraped data here. Must end with a '/' (e.g., "some_dir/")
const SCRAPES_DIR = "scraped_sites/"

func somethingBroke(err os.Error) bool {
	if err != nil {
		fmt.Printf("%s\n", err)
		return true
	}
	return false
}

// scrape scrapes the given URL and saves it to disk
func scrape(url string, done chan bool) {
	// Notify main when we're done right after we return
	defer func() { done <-true }() // Anonymous functions ftw!

	fmt.Printf("Scraping %s...\n", url)
	defer fmt.Printf("Finished scraping %s\n", url)

	// Don't make the user type "http://" for every freaking URL!
	if !strings.Contains(url, "://") {
		url = "http://" + url
	}

	// Download website contents
	req, err := http.Get(url)
	if somethingBroke(err) { return }

	// Save contents to variable
	contents, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if somethingBroke(err) { return }

	// Write contents to disk. TODO: Store URL, text data in a DB
	url = strings.Replace(url, "/", "___", -1)
	filename := fmt.Sprintf("%v-%v", url, time.Seconds())
	err = ioutil.WriteFile(SCRAPES_DIR + filename, contents, 0644)
	if somethingBroke(err) { return }

	return
}


func main() {
	os.Mkdir(SCRAPES_DIR, 0755)
	done := make(chan bool)
	urlList := []string{
		"http://git-osx-installer.googlecode.com/files/" +
			"git-1.7.7.3-intel-universal-snow-leopard.dmg",
		"steveswebapps.com/decentra/sha1/hash/" +
			"d22672d4e9406d5b74868e29b452551b637af487",
		"checkip.dyndns.org",
		"delicious.com",
		"reddit.com",
	}
	// Non-blocking scraping! Git download starts first, finishes last
	for _, url := range urlList { // Scripting language-esque for-loop
		go scrape(url, done)
	}
	// Stop main from returning till all goroutines have finished
	for i := 0; i < len(urlList); i++ { // Classic for-loop
		<-done
	}
	fmt.Printf("\nLook in %s for scraped site data\n", SCRAPES_DIR)
}

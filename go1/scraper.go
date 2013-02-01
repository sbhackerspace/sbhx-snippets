// Steve Phillips / elimisteve
// First pushed 2011.01.16
// Rewritten    2013.02.01

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type SiteScrape struct {
	URL       string     // Website URL
	Contents  []byte     // Page contents
	Timestamp time.Time  // Seconds elapsed since Unix epoch
}

func main() {
	urlList := []string{
		"onlyhdwallpapers.com/wallpaper/outer_space_galaxies_planets_desktop_" +
			"2560x1600_hd-wallpaper-1007759.jpg",
		"checkip.dyndns.org",
		"reddit.com",
		"delicious.com",
	}
	sites := make(chan *SiteScrape)

	// Non-blocking scraping! Wallpaper download starts first, finishes last
	for _, url := range urlList {
		fmt.Printf("Scraping %s...\n", url)
		go scrape(url, sites)
	}

	for i := 0; i < len(urlList); i++ {
		if site := <-sites; site != nil {
			// Display data about successful scrapes
			fmt.Printf("Scraped %s (downloaded %d bytes at %v)\n",
				site.URL, len(site.Contents),
				site.Timestamp.Format(time.Kitchen))
		}
	}
}

// scrape scrapes the given URL, puts its contents into a new
// *SiteScrape, then passes it to the `results` chan
func scrape(url string, results chan *SiteScrape) {
	var site *SiteScrape

	// Pass scraped site to `results` on success, pass `nil` on error
	defer func() {
		results <- site
	}()

	// Don't make the user type "http://" for every URL
	if !strings.Contains(url, "://") {
		url = "http://" + url
	}

	// Download website contents
	req, err := http.Get(url)
	if somethingBroke(err) {
		return
	}

	// Save contents to variable
	contents, err := ioutil.ReadAll(req.Body)
	if somethingBroke(err) {
		return
	}
	defer req.Body.Close()

	// Define new struct
	site = &SiteScrape{
		URL: url,
		Contents: contents,
		Timestamp: time.Now(),
	}

	// See `results <- site` within defer above
}

func somethingBroke(err error) bool {
	if err != nil {
		fmt.Printf("%s\n", err)
		return true
	}
	return false
}

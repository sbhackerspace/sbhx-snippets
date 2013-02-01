// Steve Phillips / elimisteve
// 2011.01.16

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
	scrapes := make(chan *SiteScrape)
	sites := []*SiteScrape{}

	// Non-blocking scraping! Wallpaper download starts first, finishes last
	fmt.Printf(" * Parallelized scraper started...\n")
	for _, url := range urlList {
		go scrape(url, scrapes)
	}

	// Collect results from scraping 
	fmt.Printf("\n * Collecting results...\n\n")
	for i := 0; i < len(urlList); i++ {
		sites = append(sites, <-scrapes)
	}

	fmt.Printf("\n * Displaying results...\n")
	for i := 0; i < len(sites); i++ {
		maxLen := 200
		if len(sites[i].Contents) < maxLen {
			maxLen = len(sites[i].Contents)
		}
		fmt.Printf("\n[%v] %s\n%s\n\n", sites[i].Timestamp.Format(time.Kitchen),
			sites[i].URL, sites[i].Contents[:maxLen])
	}
}

// scrape scrapes the given URL and saves it to disk
func scrape(url string, results chan *SiteScrape) {
	var err error

	fmt.Printf("Started scraping %s...\n", url)
	defer func() {
		if err == nil {
			fmt.Printf("Successfully scraped %s\n", url)
		}
	}()

	// Don't make the user type "http://" for every freaking URL!
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
	site := SiteScrape{
		URL: url,
		Contents: contents,
		Timestamp: time.Now(),
	}

	// Pass results to `results` channel
	results <- &site

	return
}

func somethingBroke(err error) bool {
	if err != nil {
		fmt.Printf("%s\n", err)
		return true
	}
	return false
}

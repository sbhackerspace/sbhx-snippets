// Steve Phillips / elimisteve
// 2011.01.16

package main

import (
	"fmt"
	"http"
	"io/ioutil"
)

func scrape(url string) (html string) {
	req, err := http.Get("http://" + url)
	if err == nil {
		body, er := ioutil.ReadAll(req.Body)
		if er != nil {
			fmt.Printf("%v\n", er)
		}
		req.Body.Close()
		html = string(body)
	} else {
		fmt.Printf("%v\n", err)
	}
	return
}

func main() {
	urlList := []string{
		"steveswebapps.com/decentra/md5/string/hash-this-string",
		"steveswebapps.com/decentra/sha1/hash/" +
			"d22672d4e9406d5b74868e29b452551b637af487",
		"checkip.dyndns.org",
	}
	for _, url := range urlList {
		fmt.Printf("%s: %s\n", url, scrape(url))
	}
}

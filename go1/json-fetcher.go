// Steve Phillips / elimisteve
// 2014.08.08

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/mreiferson/go-httpclient"
)

const (
	HttpGetTimeout = 60 * time.Second
)

func main() {
	hyperactiveUrl := "http://tryingtobeawesome.com:9999/services"
	var services []*HypeService

	err := FetchAndJSONUnmarshal(hyperactiveUrl, &services)
	if err != nil {
		log.Fatalf("Error from FetchAndJSONUnmarshal: %s\n", err)
	}
	for _, service := range services {
		log.Printf("%#v\n", service)
	}
}

type HypeService struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

func FetchAndJSONUnmarshal(url string, structure interface{}) error {
	body, err := Fetch(url)
	if err != nil {
		return fmt.Errorf("Error fetching %s: %v", url, err)
	}
	return json.Unmarshal(body, structure)
}

func Fetch(url string) ([]byte, error) {
	log.Printf("Fetching %s\n", url)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating GET request: %v", err)
	}

	// Create an HTTP client that times out after HttpGetTimeout seconds
	transport := &httpclient.Transport{RequestTimeout: HttpGetTimeout}
	defer transport.Close()

	client := &http.Client{Transport: transport}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// log.Printf("Returning from `fetch`: %s\n", body)
	return body, nil
}

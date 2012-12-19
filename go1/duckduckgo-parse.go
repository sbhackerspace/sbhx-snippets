// Steve Phillips / elimisteve
// 2012.12.18

package main

import (
	"fmt"
	"github.com/elimisteve/ddg"
	"log"
)

const (
	DDG_API_BASE_URL   = "api.duckduckgo.com"
	DDG_API_URL        = "http://api.duckduckgo.com/?q=%s&format=json"
	DDG_API_URL_SECURE = "https://api.duckduckgo.com/?q=%s&format=json"
	DEBUG = true
)

func main() {
	query := "define wikipedia"
	resp, err := ddg.Query(query)
	if err != nil {
		log.Printf("Error calling ddg.Query: %v\n", err)
		return
	}
	fmt.Printf("%s: %s\n", query, resp.Abstract)
}
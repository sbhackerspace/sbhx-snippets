// Steve Phillips / elimisteve
// 2011.01.16

package main

import (
	"fmt"
	"http"
	"io/ioutil"
)

func main() {
	//url := "http://decentra.org/md5/string/hash-this-string"
	url := "http://steveswebapps.com/decentra/md5/string/hash-this-string"
	req, err := http.Get(url)
	if err == nil {
		body, _ := ioutil.ReadAll(req.Body)
		fmt.Printf("%s\n", body)
		req.Body.Close()
	} else {
		fmt.Printf("err: %s\n", err)
	}
}
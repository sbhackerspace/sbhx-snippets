// Steve Phillips / elimisteve
// 2012.10.27

package main

import (
	"fmt"
	"io/ioutil"
	"github.com/elimisteve/fun"
	"net/http"
)

const USER_AGENT = "Mozilla/5.0 (Windows NT 6.1; rv:10.0) Gecko/20100101 Firefox/10.0"

func main() {
	// Be sure to run 'nc -l -p 9999' or else this won't work. Install
	// netcat with 'sudo apt-get install netcat-traditional'

	req, err := http.NewRequest("GET", "http://localhost:9999", nil)
	fun.MaybeFatalAt("NewRequest", err)

	// Change User Agent String (default: 'Go http package')
	req.Header["User-Agent"] = []string{USER_AGENT}

	client := http.Client{}
	resp, err := client.Do(req)
	fun.MaybeFatalAt("client.Do", err)

	body, err := ioutil.ReadAll(resp.Body)
	fun.MaybeFatalAt("ReadAll", err)
	defer resp.Body.Close()

	fmt.Printf("%s", body)
}
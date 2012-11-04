// Steve Phillips / elimisteve
// 2012.10.27

package main

import (
	"fmt"
	"github.com/elimisteve/fun"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	PROXY_URL        = "http://localhost:8118"
	REMOTE_HTTP_ADDR = "http://ifconfig.me/ip"
)

func main() {
	// Before running this, be sure to install Polipo: 'sudo apt-get
	// install polipo'. Get the config file at
	// https://gitweb.torproject.org/torbrowser.git/blob_plain/1ffcd9dafb9dd76c3a29dd686e05a71a95599fb5:/build-scripts/config/polipo.conf
	// then save it to /etc/polipo/config and restart Polipo with
	// 'sudo service polipo restart', _then_ run this program.

	go noProxy()

	proxyURL, err := url.Parse(PROXY_URL)
	fun.MaybeFatalAt("url.Parse", err)

	transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	// fmt.Printf("transport == %+v\n", transport)
	client := &http.Client{Transport: transport}

	request, err := http.NewRequest("GET", REMOTE_HTTP_ADDR, nil)
	fun.MaybeFatalAt("NewRequest", err)

	resp, err := client.Do(request)
	fun.MaybeFatalAt("client.Do", err)

	// resp, err := http.Get(REMOTE_HTTP_ADDR)
	// fun.MaybeFatalAt("Get", err)
	body, err := ioutil.ReadAll(resp.Body)
	fun.MaybeFatalAt("ReadAll", err)
	defer resp.Body.Close()

	fmt.Printf("Torified: %s", body)
}

func noProxy() {
	resp, err := http.Get(REMOTE_HTTP_ADDR)
	fun.MaybeFatalAt("Get", err)

	body, err := ioutil.ReadAll(resp.Body)
	fun.MaybeFatalAt("ReadAll", err)
	defer resp.Body.Close()

	fmt.Printf("Direct:   %s", body)
}

// Steve Phillips / elimisteve
// 2012.10.27

package main

import (
	"code.google.com/p/go.net/proxy"
	"fmt"
	"github.com/elimisteve/fun"
)

const (
	PROXY_URL  = "127.0.0.1:9050"
	TCP_IP_URL = "checkip.dyndns.com:80"
)

func main() {
	// Create dialer that proxies through PROXY_URL
	dialer, err := proxy.SOCKS5("tcp", PROXY_URL, nil, proxy.Direct)
	fun.MaybeFatalAt("proxy.SOCKS5", err)

	// Connect to checkip.dyndns.com
	conn, err := dialer.Dial("tcp", TCP_IP_URL)
	fun.MaybeFatalAt("dialer.Dial", err)

	// Request '/'
	_, err = conn.Write([]byte("GET /\n\n"))
	fun.MaybeFatalAt("conn.Write", err)

	// Read response
	buf := make([]byte, 1e5)
	n, err := conn.Read(buf)
	fun.MaybeFatalAt("conn.Read", err)

	// Print response from checkip.dyndns.com
	fmt.Printf("%s\n", buf[:n])
}

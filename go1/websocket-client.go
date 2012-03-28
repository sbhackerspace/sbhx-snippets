// Steve Phillips / elimisteve
// 2012.03.27

// Copied from http://golang.org/pkg/websocket/

package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"log"
)

func main() {
 	ws, err := websocket.Dial("ws://localhost:12345/echo", "",
		"http://localhost/")
	if err != nil {
		log.Fatalf("Dial: %v", err)
	}
	if _, err := ws.Write([]byte("hello, world!")); err != nil {
		log.Fatalf("Write: %v", err)
	}
	msg := make([]byte, 512)
	n, err := ws.Read(msg)
	if err != nil {
		log.Fatalf("Read: %v", err)
	}
	fmt.Printf("Sent '%s' to server...\n", msg[:n])
}

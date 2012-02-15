// From http://golang.org/pkg/websocket/

package main

import (
	"fmt"
	"websocket"
	// "strings"
)

func main() {
 	ws, err := websocket.Dial("ws://localhost/ws", "", "http://localhost/");
 	if err != nil {
		panic("Dial: " + err.String())
	}
	if _, err := ws.Write([]byte("hello, world!\n")); err != nil {
		panic("Write: " + err.String())
	}
	var msg = make([]byte, 512)
	var n int
	if n, err = ws.Read(msg); err != nil { ; }
	fmt.Println(msg[:n])
}
// Steve Phillips / elimisteve
// 2012.03.27

// Copied from http://golang.org/pkg/websocket/

package main

import (
	"net/http"
	"fmt"
	"log"
	"code.google.com/p/go.net/websocket"
)

// Echo the data received on the Web Socket.
func EchoServer(ws *websocket.Conn) {
	// io.Copy(ws, ws);
	var msg = make([]byte, 512);
	n, err := ws.Read(msg)
	if err != nil {
		log.Fatalf("Read: %v", err)
	}
	reply := fmt.Sprintf("You said: '%s'", msg[:n])
	_, err = ws.Write([]byte(reply))
	if err != nil {
			log.Fatalf("Write: %v", err)
	}
	fmt.Printf("Replied to user with...\n%s\n", reply)
}

func main() {
	http.Handle("/echo", websocket.Handler(EchoServer));
	err := http.ListenAndServe(":12345", nil);
	if err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
// Steve Phillips / elimisteve
// 2012.03.27

// Copied from http://golang.org/pkg/websocket/

package main

import (
	"net/http"
	"fmt"
	"log"
	"math/rand"
	"time"
	"code.google.com/p/go.net/websocket"
)

// Echo the data received on the Web Socket.
func EchoServer(ws *websocket.Conn) {
	// io.Copy(ws, ws);

	// var msg = make([]byte, 512);
	// n, err := ws.Read(msg)
	// if err != nil {
	// 	log.Fatalf("Read: %v", err)
	// }

	var errCount int
	for errCount < 10 {
		num := fmt.Sprintf("%v", rand.Intn(10))
		_, err := ws.Write([]byte(num))
		if err != nil {
			fmt.Printf("Something went wrong! Error: %s\n", err)
			errCount++
			time.Sleep(1e9)
		}
		fmt.Printf("To the client: %s\n", num)
		time.Sleep(0.3 * 1e9)
	}
	fmt.Printf("Too many errors! Quitting.\n")
}

func main() {
	http.Handle("/echo", websocket.Handler(EchoServer));
	err := http.ListenAndServe(":8080", nil);
	if err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
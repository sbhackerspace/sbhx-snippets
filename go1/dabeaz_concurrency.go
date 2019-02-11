// Steve Phillips / elimisteve
// 2019.02.11

// Inspired by this great Python concurrency talk by David Beazley:
// https://youtu.be/MCs5OvhV9S4

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

func main() {
	port := "25000"
	log.Printf("Listening on port %v\n", port)
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error accepting: %v\n", err)
			continue
		}

		// In Go, to launch a function into a goroutine (cheap
		// thread), just put the 'go' keyword in front of it!

		fibHandler(conn) // <-- blocking version
		// go fibHandler(conn) // <-- non-blocking version(!)
	}
}

func fibHandler(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 100)

	for {
		n, err := conn.Read(buf[:])
		if err != nil {
			if err == io.EOF {
				return
			}
			conn.Write([]byte("Error: " + err.Error() + "\n"))
			return
		}

		if n <= 1 {
			return
		}

		numStr := strings.TrimRight(string(buf[:n]), "\r\n")
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Printf("Error: didn't input number: %v\n", err)
			return
		}

		answer := fib(num)
		conn.Write([]byte(fmt.Sprintf("%d\n", answer)))
	}
}

func fib(num int) int {
	if num <= 1 {
		return num
	}
	return fib(num-1) + fib(num-2)
}

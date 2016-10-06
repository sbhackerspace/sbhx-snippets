// Steve Phillips / elimisteve
// 2016.10.06

package main

import (
	"io"
	"log"
	"net"
	"os"
	"path"
	"strings"
)

var (
	CORRECT_ANSWER = ""
	REWARD         = ""
)

func init() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %v <correct_answer> <reward>", path.Base(os.Args[0]))
	}

	CORRECT_ANSWER = os.Args[1]
	REWARD = os.Args[2]
}

func main() {
	port := "9999"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

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

		go handleGuess(conn)
	}
}

func handleGuess(conn net.Conn) {
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

		guess := strings.TrimRight(string(buf[:n]), "\r\n")

		if guess == CORRECT_ANSWER {
			log.Printf("CORRECT guess: `%s`", guess)
			conn.Write([]byte(guess + " - CORRECT! Your reward: " + REWARD + "\n"))
			return
		}

		log.Printf("Incorrect guess: `%s`", guess)
		conn.Write([]byte(guess + " - Nope! Try again...\n"))
	}
}

package main

// From https://raw.github.com/mattn/go-xmpp/master/example/example.go

import (
	"bufio"
	"fmt"
	"flag"
	"github.com/mattn/go-xmpp"
	"github.com/mattn/go-iconv"
	"log"
	"os"
	"strings"
)

var server   = flag.String("server", "talk.google.com:443", "server")
var username = flag.String("username", "", "username")
var password = flag.String("password", "", "password")

func fromUTF8(s string) string {
	ic, err := iconv.Open("char", "UTF-8")
	if err != nil {
		return s
	}
	defer ic.Close()
	ret, _ := ic.Conv(s)
	return ret
}

func toUTF8(s string) string {
	ic, err := iconv.Open("UTF-8", "char")
	if err != nil {
		return s
	}
	defer ic.Close()
	ret, _ := ic.Conv(s)
	return ret
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: example [options]\n")
		flag.PrintDefaults()
		os.Exit(2)
	}
	flag.Parse()
	if *username == "" || *password == "" {
		flag.Usage()
	}

	talk, err := xmpp.NewClient(*server, *username, *password)
	if err != nil {
		log.Fatal(err)
	}

	// Receive new messages
	go func() {
		for {
			chat, err := talk.Recv()
			if err != nil {
				log.Fatal(err)
			}
			text := strings.TrimRight(chat.Text, " \t\n")
			fmt.Printf("%s %s", chat.Remote, text)
		}
	}()

	// Listen on STDIN for messages of the form 'recipient@gmail.com
	// Here is the message'
	for {
		in := bufio.NewReader(os.Stdin)
		line, err := in.ReadString('\n')
		if err != nil {
			continue
		}
		line = strings.TrimRight(line, "\n")

		tokens := strings.SplitN(line, " ", 2)
		if len(tokens) == 2 {
			talk.Send(xmpp.Chat{Remote: tokens[0], Type: "chat",
			Text: toUTF8(tokens[1])})
		}
	}
}
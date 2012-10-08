package main

// Adopted from https://raw.github.com/mattn/go-xmpp/master/example/example.go

import (
	"bufio"
	"fmt"
	"github.com/mattn/go-xmpp"
	"log"
	"os"
	"strings"
)

// Struct singleton -- w00t!
var cfg = struct {
	Server, Username, Password string
}{
	Server:   "talk.google.com:443",
	Username: "YOURUSERNAME@gmail.com",
	Password: "YOURPASSWORD",
}

func main() {
	talk, err := xmpp.NewClient(cfg.Server, cfg.Username, cfg.Password)
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
			fmt.Printf("%s %s\n", chat.Remote, text)
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
			Text: tokens[1]})
		}
	}
}
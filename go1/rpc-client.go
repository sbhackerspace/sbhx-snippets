// Steve Phillips / elimisteve
// 2012.04.18

// From http://golang.org/pkg/net/rpc/

package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	serverAddress := "127.0.0.1"
	client, err := rpc.DialHTTP("tcp", serverAddress + ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	args := &Args{7,8}
	var syncReply int
	err = client.Call("Arith.Multiply", args, &syncReply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d * %d = %d\n", args.A, args.B, syncReply)

	// Asynchronous call
	quotient := new(Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	replyCall := <-divCall.Done	// will be equal to divCall
	asyncReply := replyCall.Reply.(*Quotient)
	fmt.Printf("Arith: %d / %d = %v remainder %v\n",
		args.A, args.B, asyncReply.Quo, asyncReply.Rem)
}
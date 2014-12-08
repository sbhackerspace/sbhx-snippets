// Steve Phillips / elimisteve
// 2014.12.07

package main

import (
	"encoding/json"
	"log"
)

type WontMarshal struct {
	Ints chan int
}

type WillMarshal struct {
	ints chan int
}

func (will *WillMarshal) RecvInt() int {
	return <-will.ints
}

func (will *WillMarshal) SendInt(n int) {
	will.ints <- n
}

func main() {
	// Will marshal because all exported fields are JSON-marshalable
	will := &WillMarshal{ints: make(chan int)}
	jsonData, _ := json.Marshal(will)
	log.Printf("Successfully marshaled %T: %s\n", will, jsonData)

	// Won't marshal because `Ints` is an exported channel
	wont := &WontMarshal{Ints: make(chan int)}
	_, err := json.Marshal(wont)
	if err != nil {
		log.Fatalf("Error marshaling %T: %v\n", wont, err)
	}
}

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

func (will *WillMarshal) Ints() int {
	return <-will.ints
}

func main() {
	// Will marshal
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

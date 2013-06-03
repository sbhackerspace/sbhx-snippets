// Steve Phillips / elimisteve
// 2013.06.03

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Person struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Birthday  time.Time `json:"birthday"`
}

func main() {
	jsonData := []byte(`{"birthday": "1986-1-16"}`)
	person := Person{}
	err := json.Unmarshal(jsonData, &person)
	if err != nil {
		log.Fatalf("Error running json.Unmarshal: %v\n", err)
	}
	fmt.Printf("person == %v\n", person)
}
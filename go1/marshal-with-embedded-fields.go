// Steve Phillips / elimisteve
// 2013.02.17

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {
	timestamp := Timestamp{Created: time.Now()}
	user := User{Username: "steve", Timestamper: &timestamp}
	userJson, err := json.Marshal(&user)
	if err != nil {
		log.Fatalf("Couldn't unmarshal '%+v': %v\n", user, err)
	}
	fmt.Printf("userJson == %s\n", userJson)
}

type User struct {
	Username string `json:"username"`
	Timestamper     `json:"timestamper"`
	// The big question: what will the `Timestamper` marshal to?
	// Answer: nothing
}

type Timestamper interface {
	CreatedAt() time.Time
}

// Timestamp implements Timestamper

type Timestamp struct {
	Created time.Time `json:"created"`
}

func (ts *Timestamp) CreatedAt() time.Time {
	return ts.Created
}

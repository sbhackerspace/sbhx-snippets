// Steve Phillips / elimisteve
// 2013.05.21

package main

import (
	"encoding/json"
	"fmt"
	"github.com/elimisteve/fun"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	p := &Person{FirstName: "Abraham"}
	fmt.Printf("p == %+v\n", p)
	err := json.Unmarshal([]byte("{}"), p)
	fun.MaybeFatalAt("json.Unmarshal", err)

	fmt.Printf("p after unmarshal'ing '{}' == %+v\n", p)

	hasLast := []byte(`{"last_name": "Lincoln"}`)
	err = json.Unmarshal(hasLast, p)
	fun.MaybeFatalAt("json.Unmarshal", err)
	fmt.Printf("p after unmarshal'ing 'first + last' == %+v\n", p)
}

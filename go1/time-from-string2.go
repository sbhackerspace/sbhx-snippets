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
	Birthday time.Time `json:"real_birthday"`
	BirthdayStr string `json:"birthday"`
}

func main() {
	jsonData := []byte(`{"birthday": "1986-1-16"}`)
	person := &Person{}
	err := json.Unmarshal(jsonData, person)
	if err != nil {
		log.Fatalf("Error running json.Unmarshal: %v\n", err)
	}
	bday, err := time.Parse("2006-1-2", person.BirthdayStr)
	if err != nil {
		log.Fatalf("Error running time.Parse: %v\n", err)
	}
	person.Birthday = bday
	fmt.Printf("person (json.Unmarshal) == %+v\n", person)

	p2 := &Person{Birthday: bday}
	jsonData, err = json.Marshal(p2)
	if err != nil {
		log.Fatalf("Error running second json.Marshal: %v\n", err)
	}
	fmt.Printf("jsonData == %s\n", jsonData)
}
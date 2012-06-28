// Steve Phillips / elimisteve
// 2012.06.28

package main

import (
	"fmt"
)

type PersonStruct struct {
	FirstName, LastName string
	Age int
}

type PersonMap map[string]interface{}

func main() {
	const sentence = "By now, %s %s must be %d years old\n"

	p1 := PersonStruct{FirstName: "Abraham", LastName: "Lincoln", Age: 100}
	fmt.Printf(sentence, p1.FirstName, p1.LastName, p1.Age)

	p2 := PersonMap{"FirstName": "George", "LastName": "Washington", "Age": 150}
	fmt.Printf(sentence, p2["FirstName"], p2["LastName"], p2["Age"])
}

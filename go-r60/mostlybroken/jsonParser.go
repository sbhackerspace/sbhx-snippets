// Steve Phillips / elimisteve
// 2011.09.02

package main

import (
    "fmt"
	"json"
)

// {"pk": 4827, "model": "autocall.call", "fields": {"user": 1,
// "recip_number": "7073010751"}}

type Fields struct {
	User          int
	RecipNumber   string
}

type Call struct {
	Pk       int
	Model    string
	Fields   Fields
}

func jsonParser(jsonStr []byte) {
	var call Call
	err := json.Unmarshal(jsonStr, &call)
    if err != nil {
		fmt.Printf("%v\n", err)
        return
    }
	fmt.Printf("%v\n", call)
    return
}

func main() {
	jsonStr := `{"pk": 4827, "model": "autocall.call", "fields": {"user": 1, "recip_number": "7073010751"}}`
	jsonParser([]byte(jsonStr))
}
// Steve Phillips / elimisteve
// 2013.01.13

package main

import (
	"encoding/json"
	"fmt"
	"github.com/elimisteve/fun"
)

func main() {
	jsonStr := `{"entities":{"https://thecloakproject.tent.is":true,"https://jcook.cc":true}}`
	m := map[string]map[string]bool{}
	err := json.Unmarshal([]byte(jsonStr), &m)
	fun.MaybeFatalAt("json.Unmarshal", err)
	fmt.Printf("m == %v\n", m)
}
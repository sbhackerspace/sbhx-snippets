// Steve Phillips / elimisteve
// 2013.10.27

package main

import (
	"./types"
	"log"
)

func main() {
	p := types.Profile{Name: "Steve", Age: 27}
	if err := p.Save(); err != nil {
		log.Printf("Error saving: %v\n", err)
		return
	}
	log.Printf("Saved successfully\n")
}

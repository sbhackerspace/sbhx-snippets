// Steve Phillips / elimisteve
// 2020.11.06

// Based on the JuliaCon 2019 talk: "The Unreasonable Effectiveness of
// Multiple Dispatch"

package main

import (
	"fmt"
)

type Pet interface {
	Name() string
}

type Cat struct {
	name string
}

func (c Cat) Name() string {
	return c.name
}

type Dog struct {
	name string
}

func (d Dog) Name() string {
	return d.name
}

func encounter(a, b Pet) {
	verb := meets(a, b)
	fmt.Printf("%s meets %s and %s\n",
		a.Name(), b.Name(), verb)
}

func meets(a, b Pet) string {
	_, aIsDog := a.(Dog)
	_, aIsCat := a.(Cat)
	_, bIsDog := b.(Dog)
	_, bIsCat := b.(Cat)

	if aIsDog && bIsDog {
		return "sniffs"
	}
	if aIsDog && bIsCat {
		return "chases"
	}
	if aIsCat && bIsDog {
		return "hisses"
	}
	if aIsCat && bIsCat {
		return "slinks"
	}

	return "FALLBACK"
}

func main() {
	fido := Dog{"Fido"}
	rex := Dog{"Rex"}
	whiskers := Cat{"Whiskers"}
	spots := Cat{"Spot"}

	encounter(fido, rex)
	encounter(fido, whiskers)
	encounter(whiskers, rex)
	encounter(whiskers, spots)
}

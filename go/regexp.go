// Steve Phillips / elimisteve
// 2011.02.06

package main

import (
	"flag"
	"fmt"
	"regexp"
)

func main() {
	// Simple
	matched, _ := regexp.MatchString("[0-9]+", "(123) 456-7890")
	if matched {
		fmt.Printf("Digits found\n")
	} else {
		fmt.Printf("No digits found\n")
	}
	fmt.Printf("************\n")


	// More complex
	args := flag.Args()
	phone := ""
	for _, arg := range args {
		phone += arg
	}

	if phone == "" {
		phone = "(123) 456-7890"
	}
	//phone := "987-654-3210"

	var get_digits = regexp.MustCompile("[0-9]+")
	digits := get_digits.FindAllString(phone, -1) // Read all digits

	if digits != nil {
		answer := ""
		for i := 0; i < len(digits); i++ {
			answer += digits[i]
		}
		if len(answer) == 10 {
			fmt.Printf("Phone number found: %s\n", answer)
		} else {
			fmt.Printf("Digits found: %s\n", answer)
		}
	} else {
		fmt.Printf("No digits found\n")
	}
	fmt.Printf("************\n")


	// Better: uses _range_ instead of _for_ loop
	// Uses _digits_ from previous version
	answer := ""
	for _, s := range digits {
		answer += s
	}
	if len(answer) == 10 {
		fmt.Printf("Phone number found: %s\n", answer)
	} else {
		if digits != nil {
			fmt.Printf("Digits found: %s\n", answer)
		}
	}
}

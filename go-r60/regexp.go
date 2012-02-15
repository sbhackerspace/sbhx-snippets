// Steve Phillips / elimisteve
// 2011.02.06

package main

import (
	"flag"
	"fmt"
	"regexp"
	"url"
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

	// Finds the desired URL within the larger string
	html := `flashvars="id_video=1164189&amp;theskin=default&amp;flv_url=http%3A%2F%2Fsite189.website.com%2Fvideos%2Fflv%2F3%2Fa%2Ff%2Fwebsite.com_3af235cd0b24ef5fad91b855001e00ae.flv%3Fe%3D1326724235%26ri%3D1024%26rs%3D85%26h%3D96a70406919dca69c45c4f83ce740ff9&amp;url_bigthumb=http://img100.website.com/videos/thumbs/pic.jpg&amp;`

	// (.*?) isn't a valid regexp in Go, so following is greedy
	// get_vid_url := regexp.MustCompile("flv_url=(.*)&amp;")

	// Non-greedy trick vaguely hinted at by swtch.com/~rsc/regexp/regexp1.html
	// get_vid_url := regexp.MustCompile("flv_url=(.*)&amp;.*&amp;")

	// This gets me what I need: the video URL. Won't let me escape the `.`, hmm
	get_vid_url := regexp.MustCompile("flv_url=(.*.flv)")

	// Looks like [1] will always be what we want; [0] is too broad
	for ndx, str := range get_vid_url.FindStringSubmatch(html) {
		decodedURL, _ := url.Parse(str)
		fmt.Printf("\nFindStringSubmatch[%v]: %v\n", ndx, decodedURL.Path)
	}
}

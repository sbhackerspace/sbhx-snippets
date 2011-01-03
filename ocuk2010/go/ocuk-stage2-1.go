// Steve Phillips / elimisteve
// 2010.12.16

package main

import (
	"fmt"
	"strconv"
)

var smallNums = map[string] string { "1":"one", "2":"two",
	"3":"three", "4":"four", "5":"five", "6":"six", "7":"seven",
	"8":"eight", "9":"nine", "10":"ten", "11":"eleven", "12":"twelve",
	"13":"thirteen", "14":"fourteen", "15":"fifteen", "16":"sixteen",
	"17":"seventeen", "18":"eighteen", "19":"nineteen", "20":"twenty",
	"30":"thirty", "40":"fourty", "50":"fifty", "60":"sixty",
	"70":"seventy", "80":"eighty", "90":"ninety" }

func main() {
	numInt := 1117065 % 9999
	numStr := strconv.Itoa(numInt)
	fmt.Printf("Int: %d, String: %s\n", numInt, numStr)

	for len(numStr) > 0 {
		l := len(numStr)
		switch {
		case l == 4:
			fmt.Printf("%s thousand ", smallNums[numStr[:1]])
		case l == 3:
			fmt.Printf("%s hundred and ", smallNums[numStr[:1]])
		case l == 2:
			value, ok := smallNums[numStr]
			if ok {
				fmt.Printf("%s ", value)
			} else {
				fmt.Printf("%s ", smallNums[numStr[:1]+"0"])
			}
		case l == 1:
			if numStr != "0" {
				fmt.Printf("%s!", smallNums[numStr[:1]])
			}
		}
		numStr = numStr[1:]
		if numStr[0:] == "0" {
			numStr = numStr[1:]
		}

	}
}
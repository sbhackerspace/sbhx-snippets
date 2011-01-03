// Steve Phillips / elimisteve / fraktil
// 2010.12.28

package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

func dig(num int) (int) {
	num /= 10
	return num % 10
}

func main() {
	var num1, num2 int

	i, s, d, carry := 0, 0, 0, 0
	var digits[19] int

	if len(flag.Args()) < 2 {
//		fmt.Printf("Not enough fields\n\n should be <program> <number> <number>")
		log.Exit("Not enough fields\n\n should be <program> <number> <number>")
	}

	s, _ = strconv.Atoi(flag.Arg(0))
	d, _ = strconv.Atoi(flag.Arg(1))

	num1 = s
	num2 = d

	if num1 > 999999999 || num2 > 999999999 {
//		fmt.Printf("not valid number try one from 0-999999999")
		log.Exit("not valid number try one from 0-999999999")
	}

	for i = 0; num1 > 0 && i < 9; i++ {
		digits[i] = dig(num1)
	}

	for i = 10; num2 > 0 && i < 19; i++ {
		digits[i] = dig(num2)
	}

	for i=0; i<10; i++ {
		if digits[i] + digits[i+10] > 9 {
			carry++
		}
	}

	fmt.Printf("\n\n%d", carry)
}
// Steve Phillips / elimisteve
// 2011.01.06

package main

import (
	"fmt"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
	return 0
}

func main() {
	fmt.Printf("gcd(12, 168) = %d", gcd(12, 168))
}

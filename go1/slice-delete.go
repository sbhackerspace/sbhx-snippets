// Steve Phillips / elimisteve
// 2012.06.14

package main

import (
	"fmt"
)

func main() {
	s := []int{2,5,8,11}
	const REMOVE_ME = 11

	for i, num := range s {
		if num == REMOVE_ME {
			fmt.Printf("Removing s[%d] == %v from s\n", i, num)
			s = append(s[:i], s[i+1:]...)
		}
	}
	fmt.Printf("s == %v\n", s)
}
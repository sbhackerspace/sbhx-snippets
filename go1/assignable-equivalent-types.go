// Steve Phillips / elimisteve
// 2020.03.04

package main

import (
	"fmt"

    "./pkg1"
    "./pkg2"
)

func main() {
	user1 := pkg1.User1{}
	user2 := pkg2.User2{}

	// Doesn't work:
	//   user2 = user1
	//
	// _Does_ work:
	user2 = pkg2.User2(user1)

	fmt.Printf("%v\n", user1 == pkg1.User1(user2))
}

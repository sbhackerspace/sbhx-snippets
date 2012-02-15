// Steve Phillips / elimisteve
// 2012.01.09

package main

import "os"

func main() {
	for i := len(os.Args); i > 0; i-- {
		println(os.Args[i-1])
	}
}
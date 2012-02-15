// Steve Phillips / elimisteve
// 2012.01.09

package main

import "os"

func main() {
	for _, arg := range os.Args {
		println(arg)
	}
}
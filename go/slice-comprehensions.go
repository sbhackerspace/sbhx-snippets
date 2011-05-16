// Steve Phillips / elimisteve
// 2011.04.21

package main

import (
    "fmt"
)

func main() {
    slice := []int{}
    for i := 0; i < 10; i++ {
        if i % 2 == 0 && i != 6 {
            slice = append(slice, i*3)
        }        
    }
    fmt.Printf("%v\n", slice)

    // slice := []int{i*3 for i in range(10) if i % 2 == 0 && i != 6} // Very Python-esque

    // slice := []int{append(slice, i*3) for i := below(10) if i % 2 == 0 && i != 6} // Less "clean" but more in the spirit of Go

    alphabet := []string{}
    for c := 'a'; c <= 'z'; c++ {
        alphabet = append(alphabet, string(c))
    }
    fmt.Printf("%v\n", alphabet)
    fmt.Printf("%v\n", 'z'+1)

    //alphabet := []string{for c:= between('a', 'z'); ; string(c)}
}

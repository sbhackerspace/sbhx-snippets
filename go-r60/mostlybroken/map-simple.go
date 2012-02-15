// Steve Phillips / elimisteve
// 2011.06.01

package main

import (
    "fmt"
)

func main() {
    myMap := map[string][]int{ "a": []int{1,2,3,4,5}, "b": []int{6,7,8,9,10} }
    fmt.Printf("%v\n", myMap)

    fmt.Printf("%v\n", map[string]string { "key1": "value1", })
}
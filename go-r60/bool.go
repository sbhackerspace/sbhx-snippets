// Steve Phillips / elimisteve
// 2011.03.27

package main

import (
    "fmt"
    "strconv"
)

func main() {
    if five_as_bool := strconv.Itob(5); five_as_bool {
        fmt.Printf("True\n")
    } else {
        fmt.Printf("False\n")
    }
}
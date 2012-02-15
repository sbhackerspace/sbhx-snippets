// Steve Phillips / elimisteve
// 2011.06.06

package main

import (
    "fmt"
    "strings"
)

func main() {
    var hashTypes = []string{"md4", "md5", "sha1","sha224", "sha256",
        "sha384", "sha512", "ripemd160"}
    fmt.Printf("Hash type. Options: " + strings.Join(hashTypes, `, `) + "\n")
}
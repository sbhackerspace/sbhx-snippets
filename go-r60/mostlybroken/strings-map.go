// Steve Phillips / elimisteve
// 2011.06.08

package main

import (
    "flag"
    "fmt"
    "strings"
)

func Rot13(c int) int {
    if ('a' <= c && c <= 'm') || ('A' <= c && c <= 'M') {
        return c + 13
    }
    if ('n' <= c && c <= 'z') || ('N' <= c && c <= 'Z') {
        return c - 13
    }
    return c
}

func main() {
    str := strings.Join(flag.Args(), " ")
    fmt.Printf("%s\n", strings.Map(Rot13, str))
}
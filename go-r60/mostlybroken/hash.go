// Steve Phillips / elimisteve
// 2011.05.27

// From http://groups.google.com/group/golang-nuts/browse_thread/thread/f98dc48ee6da8e7c

package main

import (
    "crypto/md4"
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "crypto/sha512"
    "crypto/ripemd160"
    // "io"
    "fmt"
    "hash"
    // "os"
)

func main() {
    for _, h := range []hash.Hash{md4.New(), md5.New(), sha1.New(),
            sha256.New224(), sha256.New(), sha512.New384(), sha512.New(),
            ripemd160.New()} {
        fmt.Printf("%x\n\n", h.Sum())
    }
}
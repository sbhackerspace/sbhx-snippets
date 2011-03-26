// Steve Phillips / elimisteve
// 2011.03.26

package main

import (
    "crypto/tls"
    "fmt"
)

func main() {
    config := &tls.Config{nil, nil, []tls.Certificate{}, nil, nil, "google.com", false}

    // config.Rand = nil
    // config.Time = nil
    // config.Certificates = nil
    // config.RootCAs = nil
    // config.NextProtos = nil
    // config.ServerName = "google.com"
    // config.AuthenticateClient = false
    // config.CipherSuites = nil  // Docs wrong?

    conn, err := tls.Dial("tcp", "", "google.com:443", config)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        buf := make([]uint8, 100)
        fmt.Printf("Reading...\n")
        fmt.Printf("Hangs :-\\\n")
        size, err := conn.Read(buf)  // Hangs
        fmt.Printf("Done reading\n")
        if err != nil {
            fmt.Printf("Error: %v\n", err)
        } else {
            fmt.Printf("Data: %v\n", buf[:size])
        }
    }
}
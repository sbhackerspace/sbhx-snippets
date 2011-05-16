// Steve Phillips / elimisteve
// 2011.03.26

package main

import (
    "fmt"
    //"xml"
    xmlx "github.com/jteeuwen/go-pkg-xmlx"
)

func main() {
    //
    // Parse XML
    //
    doc := xmlx.New()
    if err := doc.LoadFile("note.xml"); err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        notes := doc.SelectNodes("", "note")
        for i, note := range notes {
            if note != nil {
                fmt.Printf("Note #%v: %v\n", i, note)
            }            
        }
        fmt.Printf("\n")

        from := doc.SelectNodes("", "from")
        to := doc.SelectNodes("", "to")
        bodies := doc.SelectNodes("", "body")
        for i, _ := range bodies {
            fmt.Printf("From %v to %v: %v\n", from[i], to[i], bodies[i])
        }
    }
}

// Steve Phillips / elimisteve
// 2011.03.26

package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "xml"
)

type Note struct  {
    from string
    to string
    heading string
    body string
}

//type 

func main() {
    //
    // Parse XML
    //
    contents, err := ioutil.ReadFile("note.xml")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        xml_parser := xml.NewParser(strings.NewReader(string(contents)))
        // token, err := xml_parser.Token()
        // if err != nil {
        //     fmt.Printf("Error: %v\n", err)
        // } else {
        //     fmt.Printf("token is of type %T\n", token)
        // }
        note := new(Note)
        if err := xml_parser.Unmarshal(note, nil); err != nil {
            fmt.Printf("Error: %v\n", err)
        } else {
            //fmt.Printf("Note: %v\n", *note)
            fmt.Printf("%v\n", (*note).from)
        }
    }
}

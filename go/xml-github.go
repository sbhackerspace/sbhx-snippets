// Steve Phillips / elimisteve
// 2011.03.26

package main

import (
    "fmt"
    "io/ioutil"
    // "xml"
    xmlx "github.com/jteeuwen/go-pkg-xmlx"
)

func main() {
    repo_names := []string{"sbhx-snippets", "sbhx-ircbot", "sbhx-rov",
        "sbhx-sicp"} //, "sbhx-androidapp", "sbhx-projecteuler"
    for _, repo := range repo_names[:1] {
        result, err := ioutil.ReadFile(repo)
        //defer ioutil.CloseFile(repo) // method doesn't exist
        if err != nil {
            fmt.Printf("%v\n", err)
        } else {
            //
            // Parse XML
            //
            doc := xmlx.New()
            if err = doc.LoadFile(string(result)); err != nil {
                fmt.Printf("Error: %v\n", err)
            } else {
                fmt.Printf("Root:\n")
                _ = doc.SelectNode("", "TreeRoot")
                _ = doc.SelectNodes("", "item")
            }
        }
    }
}

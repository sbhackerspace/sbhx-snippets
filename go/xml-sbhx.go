// Steve Phillips / elimisteve
// 2011.03.28

package main

import (
    "fmt"
    //"io/ioutil"
    // "xml"
    xmlx "github.com/jteeuwen/go-pkg-xmlx"
)

func main() {
    doc := xmlx.New()
    err := doc.LoadUri("https://github.com/sbhackerspace/" + "sbhx-snippets" + "/commits/master.atom")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        children := doc.SelectNode("", "").Children
        fmt.Printf("Num children: %v\n", len(children))
        for i, child := range children {
            fmt.Printf("Child %v: %v\n", i, child.Children)
            fmt.Printf("\n\n")
        }
    }
}
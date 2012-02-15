// Steve Phillips / elimisteve
// 2011.05.27

package main

import (
    "fmt"
    "http"
    "io/ioutil"
    "strings"
)

func main() {
    url := "http://127.0.0.1:9999/"
    post := "var1=value1&secondvar=value2"
    b := strings.NewReader(post)
    req, err := http.Post(url, "application/x-www-form-urlencoded", b)

    if err == nil {
        fmt.Printf("Printing body...\n")
        body, _ := ioutil.ReadAll(req.Body)
	    fmt.Printf("%s\n", body)
	    req.Body.Close()
    } else {
        fmt.Printf("Error: %v\n", err)
    }
    return
}
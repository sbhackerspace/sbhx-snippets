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
    url := "http://127.0.0.1:8080/"
    b := strings.NewReader(`{"key":"value"}`)
    //h := http.Header{}
    //h.Add("ContentLength", "10")
    req, err := http.Post(url, "application/json", b)

    if err == nil {
        body, _ := ioutil.ReadAll(req.Body)
        req.Body.Close()
        fmt.Printf("Body: %s\n", body)
    } else {
        fmt.Printf("Error: %v\n", err)
    }
    return
}

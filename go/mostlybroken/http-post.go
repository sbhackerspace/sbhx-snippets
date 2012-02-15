// Steve Phillips / elimisteve
// 2011.05.28

package main

import (
    "fmt"
    "http"
    "io/ioutil"
)

func main() {
    c := http.Client{}
    url := "http://127.0.0.1:8080/newtask"
    data := map[string]string{
        "type": "MD5",
        "hash": "5ebe2294ecd0e0f08eab7690d2a6ee69",
    }

    req, err := c.PostForm(url, data)
    
    if err == nil {
        body, _ := ioutil.ReadAll(req.Body)
	    fmt.Printf("%s\n", body)
	    req.Body.Close()
    } else {
        fmt.Printf("Error: %v\n", err)
    }
    return
}
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
    jsonStr := `{"key":"value"}`
    b := strings.NewReader(jsonStr)
    req, err := http.NewRequest("POST", url, b)
    req.ContentLength = int64(len(jsonStr))
    res, err := http.DefaultClient.Do(req)
    go req.Body.Close()

    if err == nil {
        body, _ := ioutil.ReadAll(res.Body)
        fmt.Printf("%s\n", body)
    } else {
        fmt.Printf("Error: %v\n", err)
    }
    return
}

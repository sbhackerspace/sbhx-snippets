package main

// Mostly from http://blog.golang.org/2011/01/json-and-go.html

// Run with something like
//   echo '{"pk": 1000, "favnum": 1, "key": "throwmeaway"}' | jsonArbitrary
// after compiling, or with
//   echo '{"pk": 1000, "favnum": 1, "key": "throwmeaway"}' | gorun jsonArbitrary.go
// after installing gorun with
//   goinstall launchpad.net/gorun

import ("json"; "log"; "os")

func main() {
    dec := json.NewDecoder(os.Stdin)
    enc := json.NewEncoder(os.Stdout)
    for {
        var v map[string]interface{}
        if err := dec.Decode(&v); err != nil {
            log.Println(err)
            return
        }
        for k := range v {
            if k != "pk" {
                v[k] = nil, false
            }
        }
        if err := enc.Encode(&v); err != nil {
            log.Println(err)
        }
		//fmt.Printf("%v\n", v)
    }
}
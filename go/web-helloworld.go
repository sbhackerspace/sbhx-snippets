package main

//  Install web.go with...
/*
 git clone git://github.com/hoisie/web.go.git
 cd web.go
 make && make install
*/

import (
    "web"
)

func hello(val string) string { return "hello " + val }

func main() {
    web.Get("/(.*)", hello)
    web.Run("0.0.0.0:9999")
}

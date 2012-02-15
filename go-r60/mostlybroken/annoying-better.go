// Steve Phillips / elimisteve
// 2011.07.06

package main

import "fmt"

type Foo map[string]string

type Bar struct {
    thingOne string
    thingTwo int 
}

func main() {
    x := Foo{"x": "goodbye", "y": "world"}
    //y := Bar{thingOne: "hello", thingTwo: 1}
    y := Bar{"hello", 1}
    fmt.Printf("%v, %v\n", x, y)
}
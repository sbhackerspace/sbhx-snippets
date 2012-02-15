package main

type Foo map[string]string

type Bar struct {
    thingOne string
    thingTwo int 
}

func main() {
    x := new(Foo)
    y := new(Bar)
    (*y).thingOne = "hello"
    (*y).thingTwo = 1 
    (*x)["x"] = "goodbye"
    (*x)["y"] = "world"
}
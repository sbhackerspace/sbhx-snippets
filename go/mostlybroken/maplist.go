// Steve Phillips / elimisteve
// 2011.06.01

package main

import (
    "fmt"
)

type MapList []map[string]string

func main() {
    myMapList := someFunction()
    fmt.Printf("%v\n", myMapList)
}

func someFunction() MapList {
    // Fugly. I haven't seen any "real" Go code that looks like this.
    return MapList{
        map[string]string {
            "key1": "value1",
            "key2": "value2",
        },
        map[string]string {
            "someKey":    "someValue",
            "anotherKey": "anotherValue",
        },
    }
}

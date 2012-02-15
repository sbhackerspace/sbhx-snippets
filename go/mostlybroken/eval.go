// Steve Phillips / elimisteve
// 2011.06.08

package main

import (
    "exp/eval"
    "fmt"
)

func main() {
    world := eval.NewWorld()
    code, err := world.Compile(nil, `md5.New()`)
    if err == nil{
        fmt.Printf("code == %v\n", code)
    } else{
        fmt.Printf("err == %v\n", err)
    }
    
}
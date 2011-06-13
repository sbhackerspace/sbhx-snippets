package main

import (
    "exec"
    "fmt"
    "io/ioutil"
)

// **** DOESN'T WORK ****

func main() {
    if cmd, e := exec.Run("/bin/ls", nil, nil, "",
            exec.Pipe, exec.MergeWithStdout, 1); e == nil {
        b, _ := ioutil.ReadAll(cmd.Stdout)
        println("output: " + string(b))
    } else {
        fmt.Printf("Error: %v\n", e)
    }
}
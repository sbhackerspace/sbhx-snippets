// Steve Phillips / elimisteve
// 2011.05.09

package main

import (
    "exec"
    "fmt"
    "io/ioutil"
)

func main() {
    bin, err := exec.LookPath("python")
    if cmd, e := exec.Run(bin, []string{"", "-c", `print 'hello world'`}, nil, "", exec.DevNull, exec.Pipe, exec.MergeWithStdout); e == nil && err == nil {
        b, _ := ioutil.ReadAll(cmd.Stdout)
        fmt.Printf("output: %s\n", string(b))
    } else {
        fmt.Printf("e: %v\n", e)
        fmt.Printf("err: %v\n", err)
    }
}
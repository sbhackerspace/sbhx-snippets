// Steve Phillips / elimisteve
// 2011.05.09

// Copied from http://stackoverflow.com/questions/1877045/how-do-you-get-the-output-of-a-system-command-in-go

package main

import (
    "exec"
    "fmt"
    "io/ioutil"
)

func main() {
    if cmd, e := exec.Run("ls", []string{"-l", "/home"}, nil, "/bin", exec.DevNull, exec.Pipe, exec.MergeWithStdout); e == nil {
        b, _ := ioutil.ReadAll(cmd.Stdout)
        fmt.Printf("output: %v\n", string(b))
    } else {
        fmt.Printf("error: %v\n", e)
    }
}
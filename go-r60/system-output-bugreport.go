// Steve Phillips / elimisteve
// 2011.05.09

package main

import (
    "exec"
    "fmt"
    "io/ioutil"
    "strings"
)

func main() {
    commands := []string{
        `ls /home`,
        `ls '/home'`,
        `ls -l /home`,
        `echo 'Hello, one and all!'`,
        `which python`,

        // Outputs "hello, world\n" instead of "\n"
        `python -c print 'hello, world'`,

        // Outputs "\n" instead of "hello, world\n"
        `python -c "print 'hello, world'"`,
    }

    for ndx, c := range commands {
        fmt.Printf("Command %d of %d:  %s\n", ndx+1, len(commands), c)

        // When I replace 3 with -1, the last 2 commands produce a Python error
        argv := strings.Split(c, " ", 3)

        bin, err := exec.LookPath(argv[0])
        if cmd, e := exec.Run(bin, argv, nil, "", exec.DevNull, exec.Pipe, exec.MergeWithStdout); e == nil && err == nil {
            body, _ := ioutil.ReadAll(cmd.Stdout)
            fmt.Printf("%s\n", string(body))
        } else {
            fmt.Printf("e: %v", e)
            fmt.Printf("err: %v", err)
        }
    }
}
// Steve Phillips / elimisteve
// 2011.06.03

package main

import (
    "fmt"
    "os"
    "strings"
)

// toSeconds converts a string with format hh:mm:ss to a uint with seconds
func toSeconds(str string) (sec uint, err os.Error) {
    time := strings.Split(str, ":", -3)
    h, m, s := time[0], time[1], time[2]
    
	return uint(h)*3600+uint(m)*60+uint(s), nil
}

func main() {
    uint, err := toSeconds("10:34:22")
    if err != nil {
        fmt.Printf("%v\n", uint)
    }
}
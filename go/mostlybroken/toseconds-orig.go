// Steve Phillips / elimisteve
// 2011.06.03

package main

import (
    "fmt"
    "os"
    "strconv"
)

func toSecondsCheck(e os.Error) (uint, os.Error) {
    if e != nil {
		return 0, e
	}
    return 1, e
}

// Convert a string with format hh:mm:ss to a uint with seconds
func toSeconds(str string) (sec uint, err os.Error) {
	h, err := strconv.Atoui(str[0:2])
    toSecondsCheck(err)
	m, err := strconv.Atoui(str[3:5])
    toSecondsCheck(err)
	s, err := strconv.Atoui(str[6:8])
    toSecondsCheck(err)

	return h*3600+m*60+s, nil
}

func main() {
    uint, err := toSeconds("10!!: 22")
    if err == nil {
        fmt.Printf("%v\n", uint)
    } else {
        fmt.Printf("%v\n", err)
    }
}
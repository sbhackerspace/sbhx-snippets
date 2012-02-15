// Steve Phillips / elimisteve
// 2011.06.03

package main

import (
    "fmt"
)

func toSeconds(str string) (sec uint, err os.Error) {
	h, e := strconv.Atoui(str[0:2])
	m, er := strconv.Atoui(str[3:5])
	s, err := strconv.Atoui(str[6:8])

	if e != nil || er != nil || err ! {
		return 0, err
	}
	return h*3600+m*60+s, nil
}

func main() {
    // Convert a string with format hh:mm:ss to a uint with seconds

}
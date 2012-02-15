// Steve Phillips / elimisteve
// 2012.02.15

package main

import (
	"os/exec"
	"fmt"
	"log"
)

func checkError(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func main() {
	cmd := exec.Command("ls")
	cmd.Dir = "."  // Print contents of local dir
	output, err := cmd.Output()
	checkError(err)
	fmt.Println(string(output))
}
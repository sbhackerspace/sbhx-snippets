// Steve Phillips / elimisteve
// 2012.02.15

package main

import (
	"os/exec"
	"fmt"
	"log"
	"strings"
)

func checkError(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func main() {
	// Define paths
	const REPO_BASE_PATH = "/home/steve/sbhx/"
	const REPO_NAME = "sbhx-snippets"
	const REPO_PATH = REPO_BASE_PATH + REPO_NAME
	// Show info from latest commit
	const COMMAND = "git log -1"

	// Create *Command object
	args := strings.Split(COMMAND, " ")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = REPO_PATH  // Where cmd is run from
	output, err := cmd.Output()
	checkError(err)

	// `output` now contains COMMAND output
	lines := strings.SplitN(string(output), "\n", 4)
	// commitLine := lines[0]
	authorLine := lines[1]
	// dateLine := lines[2]
	commitMsg := strings.Replace(lines[3], "\n", "", -1)[4:]

	tokens := strings.Split(authorLine[8:], " ")
	// authorEmail := tokens[len(tokens)-1]
	authorNames := tokens[:len(tokens)-1]
	// authorFirst := authorNames[0]
	author := strings.Join(authorNames, " ")

	fmt.Printf("%v\n", author)
	fmt.Printf("%v\n", commitMsg)
}
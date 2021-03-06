package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	n := Int(Line())
	// stacks := make([][]int64, n)
	for i := int64(0); i < n; i++ {
		_ = Int(Line())
		// stacks[i] = ReadInts(Line())
		// stack := ReadInts(Line())

		// fmt.Printf("%#v  \tsum: %v; len: %d\n", stacks[i], sum(stacks[i]), len(stacks[i]))
		if good(ReadInts(Line())) {
			fmt.Printf("YES\n")
		} else {
			fmt.Printf("NO\n")
		}
	}

	// fmt.Printf("%#v\n", stacks)
}

func good(nums []int64) bool {
	// fmt.Printf("\n")
	// fmt.Printf("stack == %#v\n", nums)
	length := int64(len(nums))
	if length == 1 {
		return true
	}
	if eqNotEnough(nums) {
		return false
	}
	return true
}

func eqNotEnough(nums []int64) bool {
	cumulative := nums[0]
	for i := int64(1); i < int64(len(nums)); i++ {
		cumulative += nums[i]
		triangle := (i * (i + 1)) / 2
		// fmt.Printf("cumulative == %v\n", cumulative)
		// fmt.Printf("triangle == %v\n", triangle)
		if cumulative < triangle {
			// fmt.Printf("Too small\n")
			return true
		}
	}
	return false
}

func Run(n int64) int64 {
	return n
}

func sum(nums []int64) (total int64) {
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return
}

//
// Helpers
//

// Input

var read = bufio.NewReader(os.Stdin)

func Line() string {
	s, err := read.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(s[:len(s)-1])
}

func ReadInts(line string) []int64 {
	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanWords)
	var nums []int64
	for scanner.Scan() {
		x := Int(scanner.Text())
		nums = append(nums, x)
	}
	return nums
}

// Conversions

func Int(nStr string) int64 {
	n, err := strconv.ParseInt(nStr, 10, 64)
	if err != nil {
		panic(err)
	}
	return n
}

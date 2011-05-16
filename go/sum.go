// Steve Phillips / elimisteve
// 2011.04.17

package main

// sum takes a slice of ints and returns their sum
func sum(nums []int) (total int) {
    for _, n := range nums {
        total += n
    }
    return // Notice the named return value, 'total'
}

type MySlice []int

// Adds the Sum() method to MySlice, which has a base type of []int
func (s MySlice) Sum() int {
    slice := []int(s)
    return sum(slice)
}

func main() {
    intSlice := []int{1,2,3}
    println(sum(intSlice))
    println(MySlice(intSlice).Sum())
}

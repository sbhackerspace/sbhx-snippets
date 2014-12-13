// Steve Phillips / elimisteve
// 2014.12.12

package main

import (
	"fmt"
)

type Iter struct {
	n int
}

func (iter *Iter) Get() int {
	return iter.n
}

func (iter *Iter) Set(num int) *Iter {
	iter.n = num
	return iter
}

func (iter *Iter) Next() int {
	iter.n++
	return iter.n
}

func (iter *Iter) Skip(num int) *Iter {
	for i := 0; i < num; i++ {
		iter.Next()
	}
	return iter
}

func main() {
	fmt.Printf("%v\n", new(Iter).Set(5).Skip(10).Get())
}

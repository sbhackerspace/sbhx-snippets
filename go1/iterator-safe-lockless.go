// Steve Phillips / elimisteve
// 2014.12.12

package main

import "fmt"

type Iter struct {
	n                   int
	set, increment, out chan int
	apply               chan func(int) int
}

func NewIter(start int) *Iter {
	iter := &Iter{start, make(chan int), make(chan int), make(chan int), make(chan func(int) int)}
	go iter.loop()
	return iter
}

func (iter *Iter) loop() {
	for {
		select {
		case newNum := <-iter.set:
			iter.n = newNum
		case inc := <-iter.increment:
			iter.n += inc
		case iter.out <- iter.n:
		case f := <-iter.apply:
			iter.n = f(iter.n)
		}
	}
}

func (iter *Iter) Get() int {
	return <-iter.out
}

func (iter *Iter) Set(num int) *Iter {
	iter.set <- num
	return iter
}

func (iter *Iter) Next() int {
	iter.increment <- 1
	return <-iter.out
}

func (iter *Iter) Skip(num int) *Iter {
	// Here it _is_ safe to call iter.Next(); there's simply nothing
	// to worry about
	for i := 0; i < num; i++ {
		iter.Next()
	}
	return iter
}

func (iter *Iter) Apply(f func(int) int) *Iter {
	iter.apply <- f
	return iter
}

func main() {
	double := func(n int) int { return n * 2 }
	fmt.Printf("%v\n", NewIter(5).Skip(10).Apply(double).Get())
}

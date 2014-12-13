// Steve Phillips / elimisteve
// 2014.12.12

package main

import (
	"fmt"
	"sync"
)

type Iter struct {
	n  int
	mu sync.RWMutex
}

func (iter *Iter) Get() int {
	iter.mu.RLock()
	defer iter.mu.RUnlock()

	return iter.n
}

func (iter *Iter) Set(num int) *Iter {
	iter.mu.Lock()
	defer iter.mu.Unlock()

	iter.n = num
	return iter
}

func (iter *Iter) Next() int {
	iter.mu.Lock()
	defer iter.mu.Unlock()

	iter.n++
	return iter.n
}

func (iter *Iter) Skip(num int) *Iter {
	// Must be careful NOT to lock here, as to avoid deadlocks, since
	// iter.Next() locks.)

	// iter.mu.Lock()
	// defer iter.mu.Unlock()

	for i := 0; i < num; i++ {
		iter.Next()
	}
	return iter
}

func (iter *Iter) Apply(f func(int) int) *Iter {
	iter.mu.Lock()
	defer iter.mu.Unlock()

	iter.n = f(iter.n)
	return iter
}

func main() {
	double := func(n int) int { return n * 2 }
	fmt.Printf("%v\n", new(Iter).Set(5).Skip(10).Apply(double).Get())
}

package kits

import (
	"testing"
)

func TestGoPool(t *testing.T) {
	pool := NewGoPool(2)

	var testSlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range testSlice {
		pool.Add(1)
		go func(i int) {
			t.Logf("goroutine:%d\n", i)
			pool.Done()
		}(v)
	}
	pool.Wait()
}

package mykey

import (
	"SK-builder/internal/conf"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSnowflake(t *testing.T) {
	s := new(conf.Server)
	s.Node = 1
	var m sync.Map
	sn := NewSnowNode(s)
	var num int32 = 0
	c := make(chan struct{}, 100000)
	for i := 0; i < 100000; i++ {
		c <- struct{}{}
	}
	wg := sync.WaitGroup{}
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
		Loop:
			for {
				select {
				case <-c:
					_, id := sn.GetID()
					if _, ok := m.Load(id); ok {
						fmt.Printf("t: %d\n", id)
						break
					}
					fmt.Printf("id: %d\n", id)
					m.Store(id, 1)
					atomic.AddInt32(&num, 1)
				default:
					break Loop
				}
			}
		}(i)
	}
	wg.Wait()
	close(c)
	assert.Equal(t, int32(100000), num)
}

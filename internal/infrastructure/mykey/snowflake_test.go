package mykey

import (
	"SK-builder/internal/conf"
	"fmt"
	"sync"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSnowflake(t *testing.T) {
	s := new(conf.Server)
	s.Node = 1

	var m sync.Map
	sn := NewSnowNode(s)
	num := 0
	c := make(chan struct{}, 10000)
	d := make([]int64, 0, 1)
	for i := 0; i < 10000; i++ {
		c <- struct{}{}
	}
	wg := sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(n int) {
		Loop:
			for {
				select {
				case <-c:
					_, id := sn.GetID()
					if _, ok := m.Load(id); ok {
						fmt.Printf("t: %d\n", id)
						d = append(d, id)
					}
					fmt.Printf("id: %d\n", id)
					m.Store(id, 1)
					num++
				default:
					break Loop
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	close(c)
	fmt.Printf("d: %v\n", d)
	assert.Equal(t, num, 10000)
}

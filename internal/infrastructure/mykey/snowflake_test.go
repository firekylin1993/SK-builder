package mykey

import (
	"SK-builder/internal/conf"
	"fmt"
	"sync"
	"testing"
)

func TestSnowflake(t *testing.T) {
	s := new(conf.Server)
	s.Node = 1

	var m sync.Map
	sn := NewSnowNode(s)

	num := 0

	c := make(chan struct{}, 100000)
	for i := 0; i < int(100000); i++ {
		c <- struct{}{}
	}
	wg := sync.WaitGroup{}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
		Loop:
			for {
				select {
				case <-c:
					_, id := sn.GetID()
					if _, ok := m.Load(id); ok {
						fmt.Printf("t: %d\n", id)
						t.Error("id is repeat")
					}

					fmt.Printf("id: %d\n", id)
					m.Store(id, 1)
					num++
				default:
					wg.Done()
					break Loop
				}
			}
		}()
	}
	wg.Wait()
	close(c)
	fmt.Println(num)
}

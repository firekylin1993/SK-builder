package kits

import "sync"

type goPool struct {
	queue chan int
	wg    *sync.WaitGroup
}

func NewGoPool(size int) *goPool {
	if size <= 0 {
		size = 1
	}
	return &goPool{
		queue: make(chan int, size),
		wg:    &sync.WaitGroup{},
	}
}

func (p *goPool) Add(delta int) {
	for i := 0; i < delta; i++ {
		p.queue <- 1
	}
	for i := 0; i > delta; i-- {
		<-p.queue
	}
	p.wg.Add(delta)
}

func (p *goPool) Done() {
	<-p.queue
	p.wg.Done()
}

func (p *goPool) Wait() {
	p.wg.Wait()
}

package utils

import "sync"

type Closer struct {
	wg sync.WaitGroup
	c  chan struct{}
}

func NewCloser(i int) *Closer {
	closer := &Closer{
		wg: sync.WaitGroup{},
		c:  make(chan struct{}),
	}
	closer.wg.Add(i)
	return closer
}

func (c *Closer) Close() {
	close(c.c)
	c.wg.Wait()
}

func (c *Closer) Done() {
	c.wg.Done()
}

func (c *Closer) Wait() chan struct{} {
	return c.c
}

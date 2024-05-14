package work

import (
	"sync"
)

type Work interface {
	Task()
}

type Pool struct {
	tasks chan Worker
	wg    sync.WaitGroup
}

func New(maxGoroutines int) *Pool {
	pool := Pool{
		tasks: make(chan Worker),
	}

	pool.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range pool.tasks {
				w.Task()
			}
			pool.wg.Done()
		}()
	}

	return &pool
}

func (p *Pool) Run(w Worker) {
	p.tasks <- w
}

func (p *Pool) Shutdown() {
	close(p.tasks)
	p.wg.Wait()
}

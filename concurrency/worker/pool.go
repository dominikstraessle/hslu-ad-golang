package worker

import "sync"

type Pool struct {
	taskPool chan func()
	stopped  bool
	mx       sync.RWMutex
}

func New(workers int) *Pool {
	pool := &Pool{
		taskPool: make(chan func()),
		mx:       sync.RWMutex{},
	}
	for i := 0; i < workers; i++ {
		pool.activate()
	}
	return pool
}

func (p *Pool) activate() {
	fn := func() {
		for task := range p.taskPool {
			// check if the pool has stopped
			p.mx.RLock()
			if p.stopped {
				p.mx.RUnlock()
				return
			}
			p.mx.RUnlock()

			// run the task
			task()
		}
	}

	go fn()
}

func (p *Pool) Stop() {
	p.mx.Lock()
	defer p.mx.Unlock()
	if p.stopped {
		return
	}

	p.stopped = true
	close(p.taskPool)
}

func (p *Pool) Execute(task func()) {
	p.mx.RLock()
	defer p.mx.RUnlock()
	if p.stopped {
		return
	}
	p.taskPool <- task
}

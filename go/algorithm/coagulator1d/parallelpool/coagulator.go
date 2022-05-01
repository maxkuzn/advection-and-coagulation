package parallelpool

import (
	"context"
	"sync"

	"github.com/maxkuzn/advection-and-coagulation/algorithm/coagulator1d"
	"github.com/maxkuzn/advection-and-coagulation/internal/cell"
)

const numWorkers = 8

type parallelPoolCoagulator struct {
	kernel   coagulator1d.Kernel
	timeStep cell.FloatType

	work chan func()
	done chan struct{}

	wg     sync.WaitGroup
	cancel func()
}

func New(kernel coagulator1d.Kernel, timeStep float64) *parallelPoolCoagulator {
	return &parallelPoolCoagulator{
		kernel:   kernel,
		timeStep: cell.FloatType(timeStep),
		work:     make(chan func()),
		done:     make(chan struct{}),
	}
}

func (c *parallelPoolCoagulator) Start() error {
	ctx, cancel := context.WithCancel(context.Background())
	c.cancel = cancel

	c.wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go c.runWorker(ctx)
	}

	return nil
}

func (c *parallelPoolCoagulator) Stop() error {
	c.cancel()
	c.wg.Wait()

	return nil
}

func (c *parallelPoolCoagulator) runWorker(ctx context.Context) {
	for {
		select {
		case f := <-c.work:
			f()
			c.done <- struct{}{}
		case <-ctx.Done():
			c.wg.Done()
			return
		}
	}
}

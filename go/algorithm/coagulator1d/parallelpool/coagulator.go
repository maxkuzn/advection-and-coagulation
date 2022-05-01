package parallelpool

import (
	"context"
	"sync"

	"github.com/maxkuzn/advection-and-coagulation/internal/field1d"

	"github.com/maxkuzn/advection-and-coagulation/algorithm/coagulator"
)

const numWorkers = 8

type coag struct {
	base *coagulator.Coagulator

	work chan func()
	done chan struct{}

	wg     sync.WaitGroup
	cancel func()
}

func New(base *coagulator.Coagulator) *coag {
	return &coag{
		base: base,

		work: make(chan func(), numWorkers),
		done: make(chan struct{}, numWorkers),
	}
}

func (c *coag) Start() error {
	ctx, cancel := context.WithCancel(context.Background())
	c.cancel = cancel

	c.wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go c.runWorker(ctx)
	}

	return nil
}

func (c *coag) Stop() error {
	return nil
}

func (c *coag) runWorker(ctx context.Context) {
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

func (c *coag) Process(field, buff field1d.Field) (field1d.Field, field1d.Field) {
	var wg sync.WaitGroup

	for i := 0; i < field.Len(); i++ {
		i := i

		wg.Add(1)
		c.work <- func() {
			defer wg.Done()

			c.base.Process(field.Cell(i), buff.Cell(i), field.Sizes())
		}
	}

	wg.Wait()

	return field, buff
}

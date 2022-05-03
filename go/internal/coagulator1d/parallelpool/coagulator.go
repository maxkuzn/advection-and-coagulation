package parallelpool

import (
	"context"
	"sync"

	"github.com/maxkuzn/advection-and-coagulation/algorithm/coagulation"

	"github.com/maxkuzn/advection-and-coagulation/internal/field1d"
)

const numWorkers = 15

type coagulator struct {
	base coagulation.Coagulator

	work chan func()

	wg     sync.WaitGroup
	cancel func()
}

func New(base coagulation.Coagulator) *coagulator {
	return &coagulator{
		base: base,

		work: make(chan func(), numWorkers),
	}
}

func (c *coagulator) Start() error {
	ctx, cancel := context.WithCancel(context.Background())
	c.cancel = cancel

	c.wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go c.runWorker(ctx)
	}

	return nil
}

func (c *coagulator) Stop() error {
	c.cancel()
	c.wg.Wait()

	return nil
}

func (c *coagulator) runWorker(ctx context.Context) {
	for {
		select {
		case f := <-c.work:
			f()
		case <-ctx.Done():
			c.wg.Done()
			return
		}
	}
}

func (c *coagulator) Process(field, buff field1d.Field) (field1d.Field, field1d.Field) {
	var wg sync.WaitGroup

	for i := 0; i < field.Len(); i++ {
		i := i

		wg.Add(1)
		c.work <- func() {
			defer wg.Done()

			c.base.Process(field.Cell(i), buff.Cell(i), field.Volumes())
		}
	}

	wg.Wait()

	return field, buff
}

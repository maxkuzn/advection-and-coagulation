package parallelpool

import (
	"context"
	"sync"

	"github.com/maxkuzn/advection-and-coagulation/algorithm/coagulation"

	"github.com/maxkuzn/advection-and-coagulation/internal/field1d"
)

const numWorkers = 15

type coagulator struct {
	base      coagulation.Coagulator
	batchSize int

	work chan func()

	wg     sync.WaitGroup
	cancel func()
}

func New(base coagulation.Coagulator, batchSize int) *coagulator {
	return &coagulator{
		base:      base,
		batchSize: batchSize,
		work:      make(chan func(), numWorkers),
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

	for i := 0; i < field.Len(); i += c.batchSize {
		startIdx := i
		endIdx := i + c.batchSize
		if endIdx > field.Len() {
			endIdx = field.Len()
		}

		wg.Add(1)
		c.work <- func() {
			defer wg.Done()

			for j := startIdx; j < endIdx; j++ {
				c.base.Process(field.Cell(j), buff.Cell(j), field.Volumes())
			}
		}
	}

	wg.Wait()

	return field, buff
}

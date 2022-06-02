package naiveparallel

import (
	"sync"

	"github.com/maxkuzn/advection-and-coagulation/algorithm/coagulation"

	"github.com/maxkuzn/advection-and-coagulation/internal/field1d"
)

type coagulator struct {
	base      coagulation.Coagulator
	batchSize int
}

func New(base coagulation.Coagulator, batchSize int) *coagulator {
	return &coagulator{
		base:      base,
		batchSize: batchSize,
	}
}

func (c *coagulator) Start() error {
	return nil
}

func (c *coagulator) Stop() error {
	return nil
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
		go func() {
			defer wg.Done()

			for j := startIdx; j < endIdx; j++ {
				c.base.Process(field.Cell(j), buff.Cell(j), field.Volumes())
			}
		}()
	}

	wg.Wait()

	return field, buff
}

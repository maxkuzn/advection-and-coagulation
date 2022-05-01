package naiveparallel

import (
	"sync"

	"github.com/maxkuzn/advection-and-coagulation/algorithm/coagulator"
	"github.com/maxkuzn/advection-and-coagulation/internal/field1d"
)

type coag struct {
	base *coagulator.Coagulator
}

func New(base *coagulator.Coagulator) *coag {
	return &coag{base: base}
}

func (c *coag) Start() error {
	return nil
}

func (c *coag) Stop() error {
	return nil
}

func (c *coag) Process(field, buff field1d.Field) (field1d.Field, field1d.Field) {
	var wg sync.WaitGroup

	for i := 0; i < field.Len(); i++ {
		i := i

		wg.Add(1)
		go func() {
			defer wg.Done()

			c.base.Process(field.Cell(i), buff.Cell(i), field.Sizes())
		}()
	}

	wg.Wait()

	return field, buff
}

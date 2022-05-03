package sequential

import (
	"github.com/maxkuzn/advection-and-coagulation/algorithm/coagulation"
	"github.com/maxkuzn/advection-and-coagulation/internal/field1d"
)

type coagulator struct {
	base coagulation.Coagulator
}

func New(base coagulation.Coagulator) *coagulator {
	return &coagulator{base: base}
}

func (c *coagulator) Start() error {
	return nil
}

func (c *coagulator) Stop() error {
	return nil
}

func (c *coagulator) Process(field, buff field1d.Field) (field1d.Field, field1d.Field) {
	for i := 0; i < field.Len(); i++ {
		c.base.Process(field.Cell(i), buff.Cell(i), field.Volumes())
	}

	return field, buff
}

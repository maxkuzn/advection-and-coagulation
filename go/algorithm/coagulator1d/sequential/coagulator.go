package sequential

import (
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
	for i := 0; i < field.Len(); i++ {
		c.base.Process(field.Cell(i), buff.Cell(i), field.Sizes())
	}

	return field, buff
}

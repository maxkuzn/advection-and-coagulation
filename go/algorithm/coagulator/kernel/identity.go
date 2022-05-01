package kernel

import (
	"github.com/maxkuzn/advection-and-coagulation/algorithm/coagulator"
	"github.com/maxkuzn/advection-and-coagulation/internal/cell"
)

var _ coagulator.Kernel = (*identity)(nil)

type identity struct{}

func NewIdentity() *identity {
	return &identity{}
}

func (k *identity) Compute(x, y float64) cell.FloatType {
	_ = x // unused
	_ = y // unused
	return 0.5
}

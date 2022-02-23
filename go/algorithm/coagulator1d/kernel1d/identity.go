package kernel1d

import (
	"github.com/maxkuzn/advection-and-coagulation/algorithm/coagulator1d"
	"github.com/maxkuzn/advection-and-coagulation/internal/cell"
)

var _ coagulator1d.Kernel = (*identity)(nil)

type identity struct{}

func NewIdentity() *identity {
	return &identity{}
}

func (k *identity) Compute(x, y int) cell.FloatType {
	_ = x // unused
	_ = y // unused
	return 1
}

package kernel

import (
	"github.com/maxkuzn/advection-and-coagulation/algorithm/coagulation"
	"github.com/maxkuzn/advection-and-coagulation/internal/cell"
)

var _ coagulation.Kernel = (*identity)(nil)

type identity struct{}

func NewIdentity() *identity {
	return &identity{}
}

func (k *identity) Compute(x, y float64) cell.FloatType {
	_ = x // unused
	_ = y // unused
	return 1
}

func (k *identity) Len() int {
	return 1
}

func (k *identity) ComputeSubSum(rank, arg int, x float64) cell.FloatType {
	return 1
}

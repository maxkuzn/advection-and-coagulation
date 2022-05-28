package kernel

import (
	"github.com/maxkuzn/advection-and-coagulation/algorithm/coagulation"
)

var _ coagulation.Kernel = (*identity)(nil)

type identity struct{}

func NewIdentity() *identity {
	return &identity{}
}

func (k *identity) Compute(x, y float64) float64 {
	_ = x // unused
	_ = y // unused
	return 1
}

func (k *identity) Len() int {
	return 1
}

func (k *identity) ComputeSubSum(rank, arg int, x float64) float64 {
	return 1
}

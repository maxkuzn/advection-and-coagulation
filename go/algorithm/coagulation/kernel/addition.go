package kernel

import (
	"github.com/maxkuzn/advection-and-coagulation/algorithm/coagulation"
)

var _ coagulation.Kernel = (*addition)(nil)

type addition struct{}

func NewAddition() *addition {
	return &addition{}
}

func (k *addition) Compute(x, y float64) float64 {
	return x + y
}

func (k *addition) Len() int {
	return 2
}

// K(v, u) = v + u = v*1 + 1*u.
func (k *addition) ComputeSubSum(rank, arg int, x float64) float64 {
	if rank != arg {
		return 1
	}

	return x
}

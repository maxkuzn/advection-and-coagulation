package kernel

import (
	"github.com/maxkuzn/advection-and-coagulation/algorithm/coagulation"
	"github.com/maxkuzn/advection-and-coagulation/internal/cell"
)

var _ coagulation.Kernel = (*multiplication)(nil)

type multiplication struct{}

func NewMultiplication() *multiplication {
	return &multiplication{}
}

func (k *multiplication) Compute(x, y float64) cell.FloatType {
	return cell.FloatType(x * y)
}

func (k *multiplication) Len() int {
	return 1
}

func (k *multiplication) ComputeSubSum(rank, arg int, x float64) cell.FloatType {
	return cell.FloatType(x)
}

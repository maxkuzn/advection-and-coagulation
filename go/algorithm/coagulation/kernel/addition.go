package kernel

import (
	"github.com/maxkuzn/advection-and-coagulation/algorithm/coagulation"
	"github.com/maxkuzn/advection-and-coagulation/internal/cell"
)

var _ coagulation.Kernel = (*addition)(nil)

type addition struct{}

func NewAddition() *addition {
	return &addition{}
}

func (k *addition) Compute(x, y float64) cell.FloatType {
	return cell.FloatType(x + y)
}

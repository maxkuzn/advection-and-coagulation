package coagulation

import "github.com/maxkuzn/advection-and-coagulation/internal/cell"

type Kernel interface {
	Compute(x, y float64) cell.FloatType
	Len() int
	ComputeSubSum(rank, arg int, x float64) cell.FloatType
}

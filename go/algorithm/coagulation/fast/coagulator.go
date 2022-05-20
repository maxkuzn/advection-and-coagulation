package fast

import (
	"github.com/maxkuzn/advection-and-coagulation/algorithm/coagulation"
	"github.com/maxkuzn/advection-and-coagulation/internal/cell"
)

type coagulator struct {
	kernel   coagulation.Kernel
	timeStep cell.FloatType
}

func New(kernel coagulation.Kernel, timeStep float64) *coagulator {
	return &coagulator{
		kernel:   kernel,
		timeStep: cell.FloatType(timeStep),
	}
}

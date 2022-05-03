package predcorr

import (
	"github.com/maxkuzn/advection-and-coagulation/algorithm/coagulation"
	"github.com/maxkuzn/advection-and-coagulation/internal/cell"
)

type Coagulator struct {
	kernel   coagulation.Kernel
	timeStep cell.FloatType
}

func New(kernel coagulation.Kernel, timeStep float64) *Coagulator {
	return &Coagulator{
		kernel:   kernel,
		timeStep: cell.FloatType(timeStep),
	}
}

package predcorr

import (
	"github.com/maxkuzn/advection-and-coagulation/algorithm/coagulation"
)

type coagulator struct {
	kernel   coagulation.Kernel
	timeStep float64
}

func New(kernel coagulation.Kernel, timeStep float64) *coagulator {
	return &coagulator{
		kernel:   kernel,
		timeStep: timeStep,
	}
}

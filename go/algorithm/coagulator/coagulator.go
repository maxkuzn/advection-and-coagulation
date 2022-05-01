package coagulator

import (
	"github.com/maxkuzn/advection-and-coagulation/internal/cell"
)

type Coagulator struct {
	kernel   Kernel
	timeStep cell.FloatType
}

func New(kernel Kernel, timeStep float64) *Coagulator {
	return &Coagulator{
		kernel:   kernel,
		timeStep: cell.FloatType(timeStep),
	}
}

func (c *Coagulator) Start() error {
	return nil
}

func (c *Coagulator) Stop() error {
	return nil
}

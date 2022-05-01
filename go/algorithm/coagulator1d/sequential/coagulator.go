package sequential

import (
	"github.com/maxkuzn/advection-and-coagulation/algorithm/coagulator1d"
	"github.com/maxkuzn/advection-and-coagulation/internal/cell"
)

type sequentialCoagulator struct {
	kernel   coagulator1d.Kernel
	timeStep cell.FloatType
}

func New(kernel coagulator1d.Kernel, timeStep float64) *sequentialCoagulator {
	return &sequentialCoagulator{
		kernel:   kernel,
		timeStep: cell.FloatType(timeStep),
	}
}

func (c *sequentialCoagulator) Start() error {
	return nil
}

func (c *sequentialCoagulator) Stop() error {
	return nil
}

package naiveparallel

import (
	"github.com/maxkuzn/advection-and-coagulation/algorithm/coagulator1d"
	"github.com/maxkuzn/advection-and-coagulation/internal/cell"
)

type parallelCoagulator struct {
	kernel   coagulator1d.Kernel
	timeStep cell.FloatType
}

func New(kernel coagulator1d.Kernel, timeStep float64) *parallelCoagulator {
	return &parallelCoagulator{
		kernel:   kernel,
		timeStep: cell.FloatType(timeStep),
	}
}

func (c *parallelCoagulator) Start() error {
	return nil
}

func (c *parallelCoagulator) Stop() error {
	return nil
}

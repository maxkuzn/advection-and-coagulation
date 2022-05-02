package coagulation

import "github.com/maxkuzn/advection-and-coagulation/internal/cell"

type Kernel interface {
	Compute(x, y float64) cell.FloatType
}

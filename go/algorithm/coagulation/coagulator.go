package coagulation

import "github.com/maxkuzn/advection-and-coagulation/internal/cell"

type Coagulator interface {
	Process(cell, buff cell.Cell, volumes []float64)
}

package advector1d

import (
	"github.com/maxkuzn/advection-and-coagulation/internal/cell"
	"github.com/maxkuzn/advection-and-coagulation/internal/field1d"
)

type centralDifference struct {
	sigma cell.FloatType
}

func NewCentralDifference(sigma float64) *centralDifference {
	return &centralDifference{
		sigma: cell.FloatType(sigma),
	}
}

func (a *centralDifference) Process(field, buff field1d.Field) (field1d.Field, field1d.Field) {
	fieldIn := field
	fieldOut := buff

	l := field.Len()

	fieldOut.SetCell(
		0,
		cell.Sum3(
			1, fieldIn.Cell(0),
			-a.sigma/2, fieldIn.Cell(1),
			0, fieldIn.Cell(l-1),
		),
	)

	fieldOut.SetCell(
		l-1,
		cell.Sum3(
			1, fieldIn.Cell(l-1),
			0, fieldIn.Cell(0),
			a.sigma/2, fieldIn.Cell(l-2),
		),
	)

	for x := 1; x < l-1; x++ {
		fieldOut.SetCell(
			x,
			cell.Sum3(
				1, fieldIn.Cell(x),
				-a.sigma/2, fieldIn.Cell(x+1),
				a.sigma/2, fieldIn.Cell(x-1),
			),
		)
	}

	return fieldOut, fieldIn
}

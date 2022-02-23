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

	l := len(field)

	fieldOut[0] = cell.Sum3(
		1, fieldIn[0],
		-a.sigma/2, fieldIn[1],
		a.sigma/2, fieldIn[l-1],
	)

	fieldOut[l-1] = cell.Sum3(
		1, fieldIn[l-1],
		-a.sigma/2, fieldIn[0],
		a.sigma/2, fieldIn[l-2],
	)

	for x := 1; x < l-1; x++ {
		fieldOut[x] = cell.Sum3(
			1, fieldIn[x],
			-a.sigma/2, fieldIn[x+1],
			a.sigma/2, fieldIn[x-1],
		)
	}

	return fieldOut, fieldIn
}

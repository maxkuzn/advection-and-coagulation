package field1d

import (
	"math"

	"gonum.org/v1/gonum/stat/distuv"

	"github.com/maxkuzn/advection-and-coagulation/internal/cell"
)

func sizeFactor(vMin, v float64) float64 {
	const sigma_2 = 0.1
	dist := distuv.Normal{
		Mu:    vMin,
		Sigma: sigma_2,
	}

	minF := dist.Prob(vMin)
	f := dist.Prob(v)

	return f / minF
}

func coordFactor(x, limit int) float64 {
	if x >= limit {
		return 0
	}

	y := 2 * math.Pi * float64(x) / float64(limit)
	return (math.Cos(y-math.Pi) + 1) / 2
}

func Init(fieldSize, particlesSizesNum int, vMin, vMax float64) Field {
	field := New(fieldSize, particlesSizesNum, vMin, vMax)

	// Fill only first 25% of the field with factor func
	limit := fieldSize / 4
	for x := 0; x < limit; x++ {
		factor := coordFactor(x, limit)

		for i := 0; i < particlesSizesNum; i++ {
			v := field.volumes[i]
			field.cells[x][i] = cell.FloatType(factor * sizeFactor(vMin, v))
		}
	}

	return field
}

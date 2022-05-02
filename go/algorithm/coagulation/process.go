package coagulation

import (
	"github.com/maxkuzn/advection-and-coagulation/internal/cell"
)

func (c *Coagulator) Process(cell, buff cell.Cell, sizes []float64) {
	for i := range cell {
		buff[i] = cell[i] + c.processSizeHalf(cell, sizes, i)
	}

	for i := range cell {
		cell[i] += c.processSizeFull(buff, sizes, i)
	}
}

func (c *Coagulator) processSizeHalf(cell cell.Cell, sizes []float64, index int) cell.FloatType {
	L1 := c.computeL1(cell, sizes, index)
	L2 := c.computeL2(cell, sizes, index)
	currValue := cell[index]

	return c.timeStep / 2 * (L1 - currValue*L2)
}

func (c *Coagulator) processSizeFull(cell cell.Cell, sizes []float64, index int) cell.FloatType {
	L1 := c.computeL1(cell, sizes, index)
	L2 := c.computeL2(cell, sizes, index)
	currValue := cell[index]

	return c.timeStep * (L1 - currValue*L2)
}

func (c *Coagulator) computeL1(cel cell.Cell, sizes []float64, index int) cell.FloatType {
	if index == 0 {
		return 0
	}

	var res cell.FloatType
	for i := 0; i <= index; i++ {
		idx1 := i
		size1 := sizes[idx1]

		idx2 := index - i
		size2 := sizes[idx2]

		add := c.kernel.Compute(size1, size2) * cel[idx1] * cel[idx2]
		if idx1 == 0 || idx2 == 0 {
			add /= 2
		}
		res += add
	}

	gridStep := (sizes[len(sizes)-1] - sizes[0]) / float64(len(sizes)-1)
	res *= cell.FloatType(gridStep)
	return res
}

func (c *Coagulator) computeL2(cel cell.Cell, sizes []float64, index int) cell.FloatType {
	var res cell.FloatType
	for i := 0; i < len(cel); i++ {
		add := c.kernel.Compute(sizes[index], sizes[i]) * cel[i]

		if i == 0 || i+1 == len(cel) {
			add /= 2
		}

		res += add
	}

	gridStep := (sizes[len(sizes)-1] - sizes[0]) / float64(len(sizes)-1)
	res *= cell.FloatType(gridStep)
	return res
}

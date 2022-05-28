package predcorr

import (
	"github.com/maxkuzn/advection-and-coagulation/internal/cell"
)

func (c *coagulator) Process(cell, buff cell.Cell, volumes []float64) {
	for i := range cell {
		buff[i] = cell[i] + c.processSizeHalf(cell, volumes, i)
	}

	for i := range cell {
		cell[i] += c.processSizeFull(buff, volumes, i)
	}
}

func (c *coagulator) processSizeHalf(cell cell.Cell, volumes []float64, index int) float64 {
	L1 := c.computeL1(cell, volumes, index)
	L2 := c.computeL2(cell, volumes, index)
	currValue := cell[index]

	return c.timeStep / 2 * (L1 - currValue*L2)
}

func (c *coagulator) processSizeFull(cell cell.Cell, volumes []float64, index int) float64 {
	L1 := c.computeL1(cell, volumes, index)
	L2 := c.computeL2(cell, volumes, index)
	currValue := cell[index]

	return c.timeStep * (L1 - currValue*L2)
}

func (c *coagulator) computeL1(cel cell.Cell, volumes []float64, index int) float64 {
	if index == 0 {
		return 0
	}

	var res float64
	for i := 0; i <= index; i++ {
		idx1 := i
		v1 := volumes[idx1]

		idx2 := index - i
		v2 := volumes[idx2]

		add := c.kernel.Compute(v1, v2) * cel[idx1] * cel[idx2]
		if idx1 == 0 || idx2 == 0 {
			add /= 2
		}
		res += add
	}

	gridStep := (volumes[len(volumes)-1] - volumes[0]) / float64(len(volumes)-1)
	res *= gridStep
	return res
}

func (c *coagulator) computeL2(cel cell.Cell, volumes []float64, index int) float64 {
	var res float64
	for i := 0; i < len(cel); i++ {
		add := c.kernel.Compute(volumes[index], volumes[i]) * cel[i]

		if i == 0 || i+1 == len(cel) {
			add /= 2
		}

		res += add
	}

	gridStep := (volumes[len(volumes)-1] - volumes[0]) / float64(len(volumes)-1)
	res *= gridStep
	return res
}

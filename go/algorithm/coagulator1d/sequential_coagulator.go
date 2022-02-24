package coagulator1d

import (
	"github.com/maxkuzn/advection-and-coagulation/internal/cell"
	"github.com/maxkuzn/advection-and-coagulation/internal/field1d"
)

type sequentialCoagulator struct {
	kernel   Kernel
	timeStep cell.FloatType
}

func NewSequential(kernel Kernel, timeStep float64) *sequentialCoagulator {
	return &sequentialCoagulator{
		kernel:   kernel,
		timeStep: cell.FloatType(timeStep),
	}
}

func (c *sequentialCoagulator) Process(field, buff field1d.Field) (field1d.Field, field1d.Field) {
	for i := range field {
		c.processCell(field[i], buff[i])
	}

	return field, buff
}

func (c *sequentialCoagulator) processCell(cell, buff cell.Cell) {
	for i := range cell {
		buff[i] = c.processSizeHalf(cell, i)
	}

	for i := range cell {
		cell[i] = c.processSizeFull(buff, i)
	}
}

func (c *sequentialCoagulator) processSizeHalf(cell cell.Cell, index int) cell.FloatType {
	L1 := c.computeL1(cell, index)
	L2 := c.computeL2(cell, index)
	currValue := cell[index]

	return c.timeStep/2*(L1-currValue*L2) + currValue
}

func (c *sequentialCoagulator) processSizeFull(cell cell.Cell, index int) cell.FloatType {
	L1 := c.computeL1(cell, index)
	L2 := c.computeL2(cell, index)
	currValue := cell[index]

	return c.timeStep*(L1-currValue*L2) + currValue
}

func (c *sequentialCoagulator) computeL1(cel cell.Cell, index int) cell.FloatType {
	if index == 0 {
		return 0
	}

	var res cell.FloatType
	for i := 0; i <= index; i++ {
		size1 := i
		size2 := index - i

		add := c.kernel.Compute(size1, size2) * cel[size1] * cel[size2]
		if size1 == 0 || size2 == 0 {
			add /= 2
		}
		res += add
	}

	res *= c.timeStep / 2
	return res
}

func (c *sequentialCoagulator) computeL2(cel cell.Cell, index int) cell.FloatType {
	if index == 0 {
		return 0
	}

	var res cell.FloatType
	for i := 0; i <= index; i++ {
		size1 := i
		size2 := index - i

		add := c.kernel.Compute(size1, size2) * cel[size1] * cel[size2]
		if size1 == 0 || size2 == 0 {
			add /= 2
		}
		res += add
	}

	res *= c.timeStep / 2
	return res
}

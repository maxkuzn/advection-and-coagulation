package parallelpool

import (
	"github.com/maxkuzn/advection-and-coagulation/internal/cell"
	"github.com/maxkuzn/advection-and-coagulation/internal/field1d"
)

func (c *parallelPoolCoagulator) Process(field, buff field1d.Field) (field1d.Field, field1d.Field) {
	for i := 0; i < field.Len(); i++ {
		i := i

		c.work <- func() {
			c.processCell(field.Cell(i), buff.Cell(i), field.Sizes())
		}
	}

	return field, buff
}

func (c *parallelPoolCoagulator) processCell(cell, buff cell.Cell, sizes []float64) {
	for i := range cell {
		buff[i] = c.processSizeHalf(cell, sizes, i)
	}

	for i := range cell {
		cell[i] = c.processSizeFull(buff, sizes, i)
	}
}

func (c *parallelPoolCoagulator) processSizeHalf(cell cell.Cell, sizes []float64, index int) cell.FloatType {
	L1 := c.computeL1(cell, sizes, index)
	L2 := c.computeL2(cell, sizes, index)
	currValue := cell[index]

	return c.timeStep/2*(L1-currValue*L2) + currValue
}

func (c *parallelPoolCoagulator) processSizeFull(cell cell.Cell, sizes []float64, index int) cell.FloatType {
	L1 := c.computeL1(cell, sizes, index)
	L2 := c.computeL2(cell, sizes, index)
	currValue := cell[index]

	return c.timeStep*(L1-currValue*L2) + currValue
}

func (c *parallelPoolCoagulator) computeL1(cel cell.Cell, sizes []float64, index int) cell.FloatType {
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
		if size1 == 0 || size2 == 0 {
			add /= 2
		}
		res += add
	}

	res *= c.timeStep / 2
	return res
}

func (c *parallelPoolCoagulator) computeL2(cel cell.Cell, sizes []float64, index int) cell.FloatType {
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
		if size1 == 0 || size2 == 0 {
			add /= 2
		}
		res += add
	}

	res *= c.timeStep / 2
	return res
}

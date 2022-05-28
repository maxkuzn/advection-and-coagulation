package cell

type Cell []float64

func New(particlesSizesNum int) Cell {
	return make([]float64, particlesSizesNum)
}

func Sum2(coef1 float64, cell1 Cell, coef2 float64, cell2 Cell) Cell {
	if len(cell1) != len(cell2) {
		panic("different len")
	}

	c := make([]float64, 0, len(cell1))

	for i := range cell1 {
		c = append(c,
			coef1*cell1[i]+
				coef2*cell2[i],
		)
	}

	return c
}

func Sum3(coef1 float64, cell1 Cell, coef2 float64, cell2 Cell, coef3 float64, cell3 Cell) Cell {
	if len(cell1) != len(cell2) || len(cell1) != len(cell3) {
		panic("different len")
	}

	c := make([]float64, 0, len(cell1))

	for i := range cell1 {
		c = append(c,
			coef1*cell1[i]+
				coef2*cell2[i]+
				coef3*cell3[i],
		)
	}

	return c
}

package cell

type FloatType float64

type Cell []FloatType

func New(particlesSizesNum int) Cell {
	return make([]FloatType, particlesSizesNum)
}

func Sum2(coef1 FloatType, cell1 Cell, coef2 FloatType, cell2 Cell) Cell {
	if len(cell1) != len(cell2) {
		panic("different len")
	}

	c := make([]FloatType, 0, len(cell1))

	for i := range cell1 {
		c = append(c,
			coef1*cell1[i]+
				coef2*cell2[i],
		)
	}

	return c
}

func Sum3(coef1 FloatType, cell1 Cell, coef2 FloatType, cell2 Cell, coef3 FloatType, cell3 Cell) Cell {
	if len(cell1) != len(cell2) || len(cell1) != len(cell3) {
		panic("different len")
	}

	c := make([]FloatType, 0, len(cell1))

	for i := range cell1 {
		c = append(c,
			coef1*cell1[i]+
				coef2*cell2[i]+
				coef3*cell3[i],
		)
	}

	return c
}

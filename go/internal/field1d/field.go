package field1d

import "github.com/maxkuzn/advection-and-coagulation/internal/cell"

type Field []cell.Cell

func New(fieldSize, particlesSizesNum int) Field {
	f := make([]cell.Cell, 0, fieldSize)

	for i := 0; i < fieldSize; i++ {
		f = append(f, cell.New(particlesSizesNum))
	}

	return f
}

package field1d

import "github.com/maxkuzn/advection-and-coagulation/internal/cell"

type Field struct {
	cells   []cell.Cell
	volumes []float64
}

func New(fieldSize, particlesSizesNum int, vMin, vMax float64) Field {
	cells := make([]cell.Cell, 0, fieldSize)

	for i := 0; i < fieldSize; i++ {
		cells = append(cells, cell.New(particlesSizesNum))
	}

	volumes := make([]float64, particlesSizesNum)
	for i := range volumes {
		volumes[i] = vMin + (vMax-vMin)*float64(i)/float64(len(cells)-1)
	}

	return Field{
		cells:   cells,
		volumes: volumes,
	}
}

func (f *Field) Len() int {
	return len(f.cells)
}

func (f *Field) Cell(i int) cell.Cell {
	return f.cells[i]
}

func (f *Field) SetCell(i int, c cell.Cell) {
	f.cells[i] = c
}

func (f *Field) Volumes() []float64 {
	return f.volumes
}

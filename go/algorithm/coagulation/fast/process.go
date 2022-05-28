package fast

import (
	"github.com/maxkuzn/advection-and-coagulation/internal/cell"
	"gonum.org/v1/gonum/dsp/fourier"
	"gonum.org/v1/gonum/mat"
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
	for a := 0; a < c.kernel.Len(); a++ {
		res += c.computeL1Rank(cel, volumes, index, a)
	}

	return res
}

func (c *coagulator) computeL1Rank(cel cell.Cell, volumes []float64, index, kernelRank int) float64 {
	g := c.kernelXcell2vec(cel, volumes, 0, kernelRank)
	t := c.kernelXcell2vec(cel, volumes, 1, kernelRank)

	fft := fourier.NewFFT(len(g))
	fftG := fft.Coefficients(nil, g)
	fftT := fft.Coefficients(nil, t)

	fftM := make([]complex128, 0, len(fftT))
	for i := range fftG {
		fftM[i] = fftG[i] * fftT[i]
	}

	mRaw := fft.Sequence(nil, fftM)

	m := mat.NewVecDense(len(cel), mRaw[:len(cel)])
	m.ScaleVec(1/float64(len(mRaw)), m)

	return float64(m.AtVec(index))
}

func (c *coagulator) kernelXcell2vec(cel cell.Cell, volumes []float64, kernelArg, kernelRank int) []float64 {
	v := make([]float64, 2*len(cel)-1)
	for i, x := range cel {
		v[i] = x * c.kernel.ComputeSubSum(kernelRank, kernelArg, volumes[i])
	}

	return v
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

package fast

import (
	"github.com/maxkuzn/advection-and-coagulation/algorithm/toeplitz"
	"github.com/maxkuzn/advection-and-coagulation/internal/cell"
	"gonum.org/v1/gonum/mat"
)

func (c *coagulator) Process(cel, buff cell.Cell, volumes []float64) {
	cellVec := mat.NewVecDense(len(cel), cel)
	buffMat := mat.NewVecDense(len(buff), buff)

	buffMat.AddVec(cellVec, c.processSizeHalf(cellVec, volumes))
	cellVec.AddVec(cellVec, c.processSizeFull(buffMat, volumes))
}

func (c *coagulator) processSizeHalf(cellVec *mat.VecDense, volumes []float64) *mat.VecDense {
	L1 := c.computeL1(cellVec, volumes)
	L2 := c.computeL2(cellVec)

	res := mat.NewVecDense(cellVec.Len(), nil)

	// c.timeStep / 2 * (L1 - currValue*L2)
	res.MulElemVec(cellVec, L2)
	res.SubVec(L1, res)
	res.ScaleVec(c.timeStep/2, res)

	return res
}

func (c *coagulator) processSizeFull(cellVec *mat.VecDense, volumes []float64) *mat.VecDense {
	L1 := c.computeL1(cellVec, volumes)
	L2 := c.computeL2(cellVec)

	res := mat.NewVecDense(cellVec.Len(), nil)

	// c.timeStep * (L1 - currValue*L2)
	res.MulElemVec(cellVec, L2)
	res.SubVec(L1, res)
	res.ScaleVec(c.timeStep, res)

	return res
}

func (c *coagulator) computeL1(cellVec *mat.VecDense, volumes []float64) *mat.VecDense {
	res := mat.NewVecDense(cellVec.Len(), nil)
	for a := 0; a < c.kernel.Len(); a++ {
		res.AddVec(res, c.computeL1Rank(cellVec, volumes, a))
	}

	res.SetVec(0, 0)

	return res
}

func (c *coagulator) computeL1Rank(cellVec *mat.VecDense, volumes []float64, kernelRank int) *mat.VecDense {
	f := c.kernelXcell2vec(cellVec, volumes, kernelRank, 0)
	g := c.kernelXcell2vec(cellVec, volumes, kernelRank, 1)

	Tg := toeplitz.Multiply(f, g)

	fVec := mat.NewVecDense(len(f), f)
	gVec := mat.NewVecDense(len(g), g)

	// g_0 * f
	fVec.ScaleVec(gVec.AtVec(0), fVec)
	// f_0 * g
	gVec.ScaleVec(fVec.AtVec(0), gVec)

	res := mat.NewVecDense(cellVec.Len(), nil)
	// h(Tg - 0.5 * T'g)
	res.AddVec(gVec, fVec)
	res.ScaleVec(0.5, res)
	res.SubVec(mat.NewVecDense(len(Tg), Tg), res)

	gridStep := (volumes[len(volumes)-1] - volumes[0]) / float64(len(volumes)-1)
	res.ScaleVec(gridStep, res)

	return res
}

func (c *coagulator) kernelXcell2vec(cellVec *mat.VecDense, volumes []float64, kernelRank, kernelArg int) []float64 {
	v := make([]float64, cellVec.Len())
	for i := 0; i < cellVec.Len(); i++ {
		x := cellVec.AtVec(i)
		v[i] = x * c.kernel.ComputeSubSum(kernelRank, kernelArg, volumes[i])
	}

	return v
}

func (c *coagulator) computeL2(cellVec *mat.VecDense) *mat.VecDense {
	vr, _ := c.v.Dims()
	interRes := mat.NewVecDense(vr, nil)
	interRes.MulVec(c.v, cellVec)

	res := mat.NewVecDense(cellVec.Len(), nil)
	res.MulVec(c.u, interRes)

	return res
}

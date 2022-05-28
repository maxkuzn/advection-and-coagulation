package coagulation

type Kernel interface {
	Compute(x, y float64) float64
	Len() int
	ComputeSubSum(rank, arg int, x float64) float64
}

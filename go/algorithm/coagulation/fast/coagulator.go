package fast

import (
	decompose2 "github.com/maxkuzn/advection-and-coagulation/algorithm/decompose"

	"github.com/maxkuzn/advection-and-coagulation/algorithm/coagulation"
	"gonum.org/v1/gonum/mat"
)

type coagulator struct {
	kernel   coagulation.Kernel
	timeStep float64
	volumes  []float64
	u        mat.Matrix
	v        mat.Matrix
}

func New(kernel coagulation.Kernel, timeStep float64, volumes []float64) (*coagulator, error) {
	k := constructK(kernel, volumes)
	u, v, err := decompose(k)
	if err != nil {
		return nil, err
	}

	return &coagulator{
		kernel:   kernel,
		timeStep: timeStep,
		volumes:  volumes,
		u:        u,
		v:        v,
	}, nil
}

func constructK(kernel coagulation.Kernel, volumes []float64) mat.Matrix {
	gridStep := (volumes[len(volumes)-1] - volumes[0]) / float64(len(volumes)-1)

	k := make([]float64, 0, len(volumes)*len(volumes))

	for _, v := range volumes {
		for col, u := range volumes {
			value := gridStep * float64(kernel.Compute(v, u))
			if col == 0 || col+1 == len(volumes) {
				value /= 2
			}
			k = append(k, value)
		}
	}

	return mat.NewDense(len(volumes), len(volumes), k)
}

func decompose(k mat.Matrix) (mat.Matrix, mat.Matrix, error) {
	u, s, v, err := decompose2.Decompose(k)
	if err != nil {
		return nil, nil, err
	}

	r, c := v.Dims()
	sv := mat.NewDense(c, r, nil)
	sv.Product(mat.NewDiagDense(len(s), s), v.T())

	return u, sv, nil
}

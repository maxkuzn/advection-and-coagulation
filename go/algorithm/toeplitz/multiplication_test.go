package toeplitz

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gonum.org/v1/gonum/mat"
)

func buildToeplitz(vec []float64) *mat.Dense {
	m := mat.NewDense(len(vec), len(vec), nil)

	for r := 0; r < len(vec); r++ {
		for c := 0; c < len(vec); c++ {
			if c > r {
				m.Set(r, c, 0)
			} else {
				m.Set(r, c, vec[r-c])
			}
		}
	}

	return m
}

func TestMultiply(t *testing.T) {
	testCases := []struct {
		name     string
		toeplitz []float64
		vec      []float64
	}{
		{
			name:     "size 1",
			toeplitz: []float64{1},
			vec:      []float64{1},
		},
		{
			name:     "size 2",
			toeplitz: []float64{2, 3},
			vec:      []float64{1, 1},
		},
		{
			name:     "size 4",
			toeplitz: []float64{1, 2, 3, 4},
			vec:      []float64{0.5, -1, 3, -1.5},
		},
		{
			name:     "size 10",
			toeplitz: []float64{-10, 5, 0.12, 5.12, -6.18, 0.35, -5, 1.3, -0.421, 1.24},
			vec:      []float64{1.124, -0.85, 3, -2, 0.11, -0.05, 6, -9, 0.11, 1.123},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			vec := mat.NewVecDense(len(tc.vec), tc.vec)
			want := mat.NewVecDense(len(tc.vec), nil)
			want.MulVec(buildToeplitz(tc.toeplitz), vec)

			got := Multiply(tc.toeplitz, tc.vec)
			gotVec := mat.NewVecDense(len(got), got)

			diff := mat.NewVecDense(want.Len(), nil)
			diff.SubVec(gotVec, want)

			const delta = 1e-6
			assert.InDelta(t, 0, diff.Norm(2), delta)
		})
	}
}

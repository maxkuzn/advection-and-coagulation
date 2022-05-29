package decompose

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gonum.org/v1/gonum/mat"
)

func TestDecompose(t *testing.T) {
	testCases := []struct {
		name    string
		matrix  *mat.Dense
		rank    int
		wantErr bool
	}{
		{
			name:   "simple",
			matrix: mat.NewDense(1, 1, []float64{1}),
			rank:   1,
		},
		{
			name:   "2x2",
			matrix: mat.NewDense(2, 2, []float64{1, 0, 0, 1}),
			rank:   2,
		},
		{
			name: "4x4 rank 2",
			matrix: mat.NewDense(4, 4, []float64{
				1, 1, 1, 3,
				1, 0, 1, 2,
				0, 1, 0, 1,
				0, 0, 0, 0,
			}),
			rank: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			u, s, v, err := Decompose(tc.matrix)

			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			r, _ := u.Dims()
			c, _ := v.Dims()
			res := mat.NewDense(r, c, nil)
			res.Product(u, mat.NewDiagDense(len(s), s), v.T())
			res.Sub(res, tc.matrix)

			assert.InDelta(t, 0, res.Norm(2), 1e-5)
		})
	}
}

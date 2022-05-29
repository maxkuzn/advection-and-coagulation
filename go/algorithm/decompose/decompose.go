package decompose

import (
	"errors"

	"gonum.org/v1/gonum/mat"
)

func Decompose(matrix mat.Matrix) (*mat.Dense, []float64, *mat.Dense, error) {
	u, s, v, rank, err := decomposeFull(matrix)

	ur, _ := u.Dims()
	uSlice := mat.NewDense(ur, rank, nil)
	uSlice.Copy(u)

	sSlice := s[:rank]

	vr, _ := u.Dims()
	vSlice := mat.NewDense(vr, rank, nil)
	vSlice.Copy(v)

	return uSlice, sSlice, vSlice, err
}

func decomposeFull(matrix mat.Matrix) (*mat.Dense, []float64, *mat.Dense, int, error) {
	var svd mat.SVD
	if !svd.Factorize(matrix, mat.SVDThin) {
		return nil, nil, nil, 0, errors.New("couldn't factorize")
	}

	const eps = 1e-6
	rank := svd.Rank(eps)

	var u, v mat.Dense
	svd.UTo(&u)
	svd.VTo(&v)
	s := svd.Values(nil)

	return &u, s, &v, rank, nil
}

package toeplitz

import (
	"gonum.org/v1/gonum/dsp/fourier"
)

func Multiply(toeplitz []float64, vec []float64) []float64 {
	// append zeros
	toeplitzSlice := make([]float64, 2*len(toeplitz)-1)
	copy(toeplitzSlice, toeplitz)

	vecSlice := make([]float64, 2*len(vec)-1)
	copy(vecSlice, vec)

	// perform fft
	fft := fourier.NewFFT(len(toeplitzSlice))
	fftToeplitz := fft.Coefficients(nil, toeplitzSlice)
	fftVec := fft.Coefficients(nil, vecSlice)

	// multiply under fft
	fftMult := make([]complex128, 0, len(fftVec))
	for i := range fftToeplitz {
		fftMult = append(fftMult, fftToeplitz[i]*fftVec[i]/complex(float64(len(vecSlice)), 0))
	}

	return fft.Sequence(nil, fftMult)[:len(vec)]
}

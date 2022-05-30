#include "multiplication.h"

#include "dj_fft/dj_fft.h"

namespace toeplitz {

std::vector<double> Multiply(const std::vector<double>& toeplitz, const std::vector<double>& vec) {
    std::vector<std::complex<double>> t_with_zeroes(2 * toeplitz.size() - 1, 0);
    for (size_t i = 0; i < toeplitz.size(); i++) {
        t_with_zeroes[i] = toeplitz[i];
    }

    std::vector<std::complex<double>> v_with_zeroes(2 * toeplitz.size() - 1, 0);
    for (size_t i = 0; i < vec.size(); i++) {
        v_with_zeroes[i] = vec[i];
    }

    auto fftT = dj::fft1d(t_with_zeroes, dj::fft_dir::DIR_FWD);
    auto fftV = dj::fft1d(v_with_zeroes, dj::fft_dir::DIR_FWD);

    std::vector<std::complex<double>> fftM;
    fftM.reserve(fftT.size());

    for (size_t i = 0; i < fftT.size(); i++) {
        fftM.push_back(fftT[i] * fftV[i] / std::complex<double>(t_with_zeroes.size()));
    }

    auto m = dj::fft1d(fftM, dj::fft_dir::DIR_BWD);

    std::vector<double> mReal;
    mReal.reserve(vec.size());

    for (size_t i = 0; i < vec.size(); i++) {
        mReal.push_back(m[i].real());
    }

    return vec;
}

}

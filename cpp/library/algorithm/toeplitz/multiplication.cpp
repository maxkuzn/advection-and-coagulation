#include "multiplication.h"

#include "dj_fft/dj_fft.h"

namespace toeplitz {

size_t closest_power_of_2(size_t x) {
    size_t pow = 1;
    size_t value = 2;

    while (value < x) {
        pow += 1;
        value *= 2;
    }

    return value;
}

std::vector<double> Multiply(const std::vector<double>& toeplitz, const std::vector<double>& vec) {
    size_t s = closest_power_of_2(2 * toeplitz.size() - 1);


    std::vector<std::complex<double>> t_with_zeroes(s, 0);
    for (size_t i = 0; i < toeplitz.size(); i++) {
        t_with_zeroes[i] = toeplitz[i];
    }

    std::vector<std::complex<double>> v_with_zeroes(s, 0);
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

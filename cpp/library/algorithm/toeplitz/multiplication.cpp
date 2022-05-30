#include "multiplication.h"

#include <Eigen/Dense>
#include <fftw3.h>

namespace toeplitz {

std::vector<double> Multiply(const std::vector<double>& toeplitz, const std::vector<double>& vec) {
    std::vector<double> t_with_zeroes(2 * toeplitz.size() - 1, 0);
    for (size_t i = 0; i < toeplitz.size(); i++) {
        t_with_zeroes[i] = toeplitz[i];
    }

    std::vector<double> v_with_zeroes(2 * toeplitz.size() - 1, 0);
    for (size_t i = 0; i < vec.size(); i++) {
        v_with_zeroes[i] = vec[i];
    }


    return vec;
}

}

#include "coagulator.h"

#include <Eigen/Dense>

#include "algorithm/toeplitz/multiplication.h"
#include "algorithm/decompose/decompose.h"

namespace coagulation {

void FastCoagulator::Process(Cell* cell, Cell* /*buff*/, const std::vector<double>& volumes) {
    Cell& cell_ref = *cell;

    Eigen::VectorXd c = Eigen::VectorXd::Map(cell_ref.Data().data(), cell_ref.Size());

    Eigen::VectorXd b(cell->Size());

    b = c + ProcessHalf(c, volumes);

    c += ProcessFull(b, volumes);

    for (size_t i = 0; i < cell_ref.Size(); i++) {
        cell_ref[i] = c[i];
    }
}

Eigen::VectorXd FastCoagulator::ProcessHalf(Eigen::VectorXd& cell, const std::vector<double>& volumes) {
    auto L1 = ComputeL1(cell, volumes);
    auto L2 = ComputeL2(cell, volumes);

    return time_step_ / 2 * (L1 - cell.cwiseProduct(L2));
}

Eigen::VectorXd FastCoagulator::ProcessFull(Eigen::VectorXd& cell, const std::vector<double>& volumes) {
    auto L1 = ComputeL1(cell, volumes);
    auto L2 = ComputeL2(cell, volumes);

    return time_step_ * (L1 - cell.cwiseProduct(L2));
}

Eigen::VectorXd FastCoagulator::ComputeL1(Eigen::VectorXd& cell, const std::vector<double>& volumes) {
    /*
    Eigen::VectorXd res(cell.size());
    res.setZero();

    for (size_t rank = 0; rank < kernel_->Size(); rank++) {
        res += ComputeL1Rank(cell, volumes, rank);
    }

    res[0] = 0;

    return res;
    */
    Eigen::VectorXd resVec(cell.size());

    for (size_t idx = 0; idx < size_t(cell.size()); idx++) {
        if (idx == 0) {
            resVec[idx] = 0;
        }

        double res = 0;
        for (size_t i = 0; i <= idx; i++) {
            size_t idx1 = i;
            double v1 = volumes[idx1];

            size_t idx2 = idx - i;
            double v2 = volumes[idx2];

            double add = kernel_->Compute(v1, v2) * cell[idx1] * cell[idx2];
            if (idx1 == 0 || idx2 == 0) {
                add /= 2;
            }

            res += add;
        }

        double grid_step = (volumes.back() - volumes.front()) / (volumes.size() - 1);
        res *= grid_step;

        resVec[idx] = res;
    }

    return resVec;
}

Eigen::VectorXd FastCoagulator::ComputeL1Rank(Eigen::VectorXd& cell, const std::vector<double>& volumes,
                                              size_t kernel_rank) {
    auto f = KernelXcell2vec(cell, volumes, kernel_rank, 0);
    auto g = KernelXcell2vec(cell, volumes, kernel_rank, 1);

    auto Tg = toeplitz::Multiply(f, g);

    auto fVec = Eigen::VectorXd::Map(f.data(), f.size());
    auto gVec = Eigen::VectorXd::Map(g.data(), g.size());

    fVec *= gVec[0] / 2;
    gVec *= fVec[0] / 2;

    Eigen::VectorXd res = Eigen::VectorXd::Map(Tg.data(), Tg.size());
    res -= gVec;
    res -= fVec;

    auto grid_step = (volumes.back() - volumes.front()) / (volumes.size() - 1);
    res *= grid_step;

    return res;
}

std::vector<double> FastCoagulator::KernelXcell2vec(Eigen::VectorXd& cell, const std::vector<double>& volumes,
                                                    size_t kernel_rank, size_t kernel_arg) {
    std::vector<double> v;
    v.resize(cell.size());

    for (size_t i = 0; i < v.size(); i++) {
        double x = cell[i];
        v[i] = x * kernel_->ComputeSubSum(kernel_rank, kernel_arg, volumes[i]);
    }

    return v;
}

Eigen::VectorXd FastCoagulator::ComputeL2(Eigen::VectorXd& cell, const std::vector<double>& /*volumes*/) {
    return u_ * (v_ * cell);

    /*
    Eigen::VectorXd resVec(cell.size());

    for (size_t idx = 0; idx < size_t(cell.size()); idx++) {
        double res = 0;
        for (size_t i = 0; i < size_t(cell.size()); i++) {
            double add = kernel_->Compute(volumes[idx], volumes[i]) * cell[i];

            if (i == 0 || i + 1 == size_t(cell.size())) {
                add /= 2;
            }

            res += add;
        }

        double grid_step = (volumes.back() - volumes.front()) / (volumes.size() - 1);
        res *= grid_step;

        resVec[idx] = res;
    }

    return resVec;
     */
}


namespace impl {

Eigen::MatrixXd constructK(const std::shared_ptr<coagulation::Kernel>& kernel, const std::vector<double>& volumes) {
    double grid_step = (volumes.back() - volumes.front()) / (volumes.size() - 1);

    Eigen::MatrixXd k(volumes.size(), volumes.size());

    for (size_t r = 0; r < volumes.size(); r++) {
        for (size_t c = 0; c < volumes.size(); c++) {
            double value = grid_step * kernel->Compute(volumes[r], volumes[c]);

            if (c == 0 || c + 1 == volumes.size()) {
                value /= 2;
            }

            k(r, c) = value;
        }
    }

    return k;
}

std::pair<Eigen::MatrixXd, Eigen::MatrixXd> decompose(Eigen::MatrixXd k) {
    auto [u, s, v] = decompose::Decompose(k);

    Eigen::MatrixXd sv = s.asDiagonal() * v.transpose();

    return std::make_pair(u, sv);
}

}


}  // namespace coagulation

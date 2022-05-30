#include "coagulator.h"

#include <Eigen/Dense>

namespace coagulation {

void FastCoagulator::Process(Cell* cell, Cell* /*buff*/, const std::vector<double>& volumes) {
    Cell& cell_ref = *cell;

    Eigen::VectorXd c(cell_ref.Size());
    for (size_t i = 0; i < cell_ref.Size(); i++) {
        c[i] = cell_ref[i];
    }

    Eigen::VectorXd b(cell->Size());

    for (size_t i = 0; i < size_t(c.size()); i++) {
        b[i] = c[i] + ProcessHalf(c, volumes, i);
    }

    for (size_t i = 0; i < size_t(c.size()); i++) {
        cell_ref[i] += ProcessFull(b, volumes, i);
    }

}

double FastCoagulator::ProcessHalf(Eigen::VectorXd& cell, const std::vector<double>& volumes, size_t idx) {
    auto L1 = ComputeL1(cell, volumes, idx);
    auto L2 = ComputeL2(cell, volumes, idx);
    auto curr_value = cell[idx];

    return time_step_ / 2 * (L1 - curr_value * L2);
}

double FastCoagulator::ProcessFull(Eigen::VectorXd& cell, const std::vector<double>& volumes, size_t idx) {
    auto L1 = ComputeL1(cell, volumes, idx);
    auto L2 = ComputeL2(cell, volumes, idx);
    auto curr_value = cell[idx];

    return time_step_ * (L1 - curr_value * L2);
}

double FastCoagulator::ComputeL1(Eigen::VectorXd& cell, const std::vector<double>& volumes, size_t idx) {
    if (idx == 0) {
        return 0;
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

    return res;
}

double FastCoagulator::ComputeL2(Eigen::VectorXd& cell, const std::vector<double>& volumes, size_t idx) {
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

    return res;
}

}  // namespace coagulation

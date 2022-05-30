#include "coagulator.h"

namespace coagulation {

void PredCorrCoagulator::Process(Cell* cell, Cell* buff, const std::vector<double>& volumes) {
    Cell& c = *cell;
    Cell& b = *buff;

    for (size_t i = 0; i < c.Size(); i++) {
        b[i] = c[i] + ProcessHalf(c, volumes, i);
    }

    for (size_t i = 0; i < c.Size(); i++) {
        c[i] += ProcessFull(b, volumes, i);
    }
}

double PredCorrCoagulator::ProcessHalf(Cell& cell, const std::vector<double>& volumes, size_t idx) {
    auto L1 = ComputeL1(cell, volumes, idx);
    auto L2 = ComputeL2(cell, volumes, idx);
    auto curr_value = cell[idx];

    return time_step_ / 2 * (L1 - curr_value * L2);
}

double PredCorrCoagulator::ProcessFull(Cell& cell, const std::vector<double>& volumes, size_t idx) {
    auto L1 = ComputeL1(cell, volumes, idx);
    auto L2 = ComputeL2(cell, volumes, idx);
    auto curr_value = cell[idx];

    return time_step_ * (L1 - curr_value * L2);
}

double PredCorrCoagulator::ComputeL1(Cell& cell, const std::vector<double>& volumes, size_t idx) {
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

double PredCorrCoagulator::ComputeL2(Cell& cell, const std::vector<double>& volumes, size_t idx) {
    double res = 0;
    for (size_t i = 0; i < cell.Size(); i++) {
        double add = kernel_->Compute(volumes[idx], volumes[i]) * cell[i];

        if (i == 0 || i + 1 == cell.Size()) {
            add /= 2;
        }

        res += add;
    }

    double grid_step = (volumes.back() - volumes.front()) / (volumes.size() - 1);
    res *= grid_step;

    return res;
}

}  // namespace coagulation

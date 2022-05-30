#pragma once

#include "algorithm/coagulation/coagulator.h"
#include "algorithm/coagulation/kernel.h"

#include <vector>
#include <memory>

#include <Eigen/Dense>

namespace coagulation {

class FastCoagulator : public Coagulator {
  public:
    FastCoagulator(std::shared_ptr<Kernel> kernel, double time_step)
            : kernel_(kernel), time_step_(time_step) {}

    ~FastCoagulator() = default;

    void Process(Cell* cell, Cell* buff, const std::vector<double>& volumes) override;

  private:
    double ProcessHalf(Eigen::VectorXd& cell, const std::vector<double>& volumes, size_t idx);

    double ProcessFull(Eigen::VectorXd& cell, const std::vector<double>& volumes, size_t idx);

    double ComputeL1(Eigen::VectorXd& cell, const std::vector<double>& volumes, size_t idx);

    double ComputeL2(Eigen::VectorXd& cell, const std::vector<double>& volumes, size_t idx);

    std::shared_ptr<Kernel> kernel_;
    const double time_step_;
};

};  // namespace coagulation

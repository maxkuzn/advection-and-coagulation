#pragma once

#include "algorithm/coagulation/coagulator.h"
#include "algorithm/coagulation/kernel.h"

#include <vector>
#include <memory>

#include <Eigen/Dense>

namespace coagulation {

namespace impl {

Eigen::MatrixXd constructK(const std::shared_ptr<coagulation::Kernel>& kernel, const std::vector<double>& volumes);

std::pair<Eigen::MatrixXd, Eigen::MatrixXd> decompose(Eigen::MatrixXd k);

}

class FastCoagulator : public Coagulator {
  public:
    FastCoagulator(std::shared_ptr<Kernel> kernel, double time_step, const std::vector<double>& volumes)
            : kernel_(kernel), time_step_(time_step) {
        auto k = impl::constructK(kernel, volumes);
        auto [u, v] = impl::decompose(k);
        u_ = u;
        v_ = v;
    }

    ~FastCoagulator() = default;

    void Process(Cell* cell, Cell* buff, const std::vector<double>& volumes) override;

  private:
    Eigen::VectorXd ProcessHalf(Eigen::VectorXd& cell, const std::vector<double>& volumes);

    Eigen::VectorXd ProcessFull(Eigen::VectorXd& cell, const std::vector<double>& volumes);

    Eigen::VectorXd ComputeL1(Eigen::VectorXd& cell, const std::vector<double>& volumes);

    Eigen::VectorXd ComputeL1Rank(Eigen::VectorXd& cell, const std::vector<double>& volumes, size_t kernel_rank);

    std::vector<double> KernelXcell2vec(Eigen::VectorXd& cell, const std::vector<double>& volumes,
                                        size_t kernel_rank, size_t kernel_arg);

    Eigen::VectorXd ComputeL2(Eigen::VectorXd& cell, const std::vector<double>& volumes);

    std::shared_ptr<Kernel> kernel_;
    const double time_step_;

    Eigen::MatrixXd u_, v_;
};

};  // namespace coagulation

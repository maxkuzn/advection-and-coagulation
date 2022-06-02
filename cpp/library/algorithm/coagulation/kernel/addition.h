#pragma once

#include "algorithm/coagulation/kernel.h"


namespace coagulation {

class AdditionKernel : public Kernel {
  public:
    AdditionKernel() = default;

    ~AdditionKernel() = default;

    double Compute(double x, double y) override {
        return x + y;
    }

    size_t Size() override {
        return 2;
    }

    double ComputeSubSum(size_t rank, size_t arg, double x) override {
        if (rank != arg) {
            return 1;
        }

        return x;
    }
};

};  // namespace coagulation


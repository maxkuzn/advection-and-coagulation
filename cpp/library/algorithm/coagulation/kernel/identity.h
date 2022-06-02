#pragma once

#include "algorithm/coagulation/kernel.h"


namespace coagulation {

class IdentityKernel : public Kernel {
  public:
    IdentityKernel() = default;

    ~IdentityKernel() = default;

    double Compute(double /*x*/, double /*y*/) override {
        return 1;
    }

    size_t Size() override {
        return 1;
    }

    double ComputeSubSum(size_t /*rank*/, size_t /*arg*/, double /*x*/) override {
        return 1;
    }
};

};  // namespace coagulation


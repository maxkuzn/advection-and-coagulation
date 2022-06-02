#pragma once

#include "base/cell.h"


namespace coagulation {

class Kernel {
  public:
    Kernel() = default;

    virtual ~Kernel() = default;

    virtual double Compute(double x, double y) = 0;

    virtual size_t Size() = 0;

    virtual double ComputeSubSum(size_t /*rank*/, size_t /*arg*/, double /*x*/) = 0;
};

};  // namespace coagulation

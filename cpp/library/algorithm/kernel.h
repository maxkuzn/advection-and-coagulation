#pragma once

#include "base/cell.h"


namespace coagulation {

class Kernel {
  public:
	Kernel() = default;
	virtual ~Kernel() = default;

	virtual double Compute(double x, double y) = 0;
};

};  // namespace coagulation

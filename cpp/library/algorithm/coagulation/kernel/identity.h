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
};

};  // namespace coagulation


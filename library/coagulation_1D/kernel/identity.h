#pragma once

#include "base.h"

class IdentityKernel : public BaseKernel {
  public:
	~IdentityKernel() = default;

	virtual double Compute(double /*x*/, double /*y*/) const {
		return 1;
	}
};


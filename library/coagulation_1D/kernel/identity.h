#pragma once

#include "base_kernel.h"

class IdentityKernel : public BaseKernel {
	~IdentityKernel() = default;

	virtual double Compute(double /*x*/, double /*y*/) const {
		return 1;
	}
};


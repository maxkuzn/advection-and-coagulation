#pragma once

class BaseKernel {
  public:
	BaseKernel() = default;
	virtual ~BaseKernel() = default;

	double Compute(double x, double y) const = 0;
};


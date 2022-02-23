#pragma once

class BaseKernel {
  public:
	BaseKernel() = default;
	virtual ~BaseKernel() = default;

	virtual double Compute(double x, double y) const = 0;
};

